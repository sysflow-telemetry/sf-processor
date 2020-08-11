package pipeline

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/driver/manifest"
)

const (
	// ChanSize is the default size of the go channels in the pipeline
	ChanSize = 100000
)

// setManifestInfo sets manifest attributes to plugins configuration items.
func setManifestInfo(conf *Config) {
	addGlobalConfigItem(conf, manifest.VersionKey, manifest.Version)
	addGlobalConfigItem(conf, manifest.JSONSchemaVersionKey, manifest.JSONSchemaVersion)
	addGlobalConfigItem(conf, manifest.BuildNumberKey, manifest.BuildNumber)
}

// addGlobalConfigItem adds a config item to all processors in the pipeline.
func addGlobalConfigItem(conf *Config, k string, v interface{}) {
	for _, c := range conf.Pipeline {
		if _, ok := c[ProcConfig]; ok {
			if s, ok := v.(string); ok {
				c[k] = s
			} else if i, ok := v.(int); ok {
				c[k] = strconv.Itoa(i)
			}
		}
	}
}

// LoadPipeline sets up the an edge processing pipeline based on configuration settings.
func LoadPipeline(pluginDir string, config string) (interface{}, []plugins.SFProcessor, *sync.WaitGroup, []interface{}, []plugins.SFHandler, error) {
	pl := NewPluginCache(config)
	wg := new(sync.WaitGroup)

	var processors []plugins.SFProcessor
	var channels []interface{}
	var hdlrs []plugins.SFHandler

	if err := pl.LoadPlugins(pluginDir); err != nil {
		logger.Error.Println("Unable to load dynamic plugins: ", err)
		return nil, nil, wg, nil, nil, err
	}

	conf, err := pl.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return nil, nil, wg, nil, nil, err
	}
	setManifestInfo(conf)
	var in interface{}
	var out interface{}
	var first interface{}
	for idx, p := range conf.Pipeline {
		hdler := false
		var hdl plugins.SFHandler
		if val, ok := p[HdlConfig]; ok {
			hdl, err = pl.GetHandler(val)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
			hdlrs = append(hdlrs, hdl)
			xType := fmt.Sprintf("%T", hdl)
			logger.Trace.Println(xType)
			hdler = true
		}
		var prc plugins.SFProcessor
		if val, ok := p[ProcConfig]; ok {
			prc, err = pl.GetProcessor(val, hdl, hdler)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
			tp := fmt.Sprintf("%T", prc)
			logger.Trace.Println(tp)
			err = prc.Init(p)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil, wg, nil, nil, err
			}
		} else {
			logger.Error.Println("processor or handler tag must exist in plugin config")
			return nil, nil, wg, nil, nil, err
		}
		if v, o := p[InChanConfig]; o {
			in, err = pl.GetChan(v, ChanSize)
			channels = append(channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return nil, nil, wg, nil, nil, errors.New("in tag must exist in plugin config")
		}
		if v, o := p[OutChanConfig]; o {
			out, err = pl.GetChan(v, ChanSize)
			chp := fmt.Sprintf("%T", out)
			channels = append(channels, out)
			logger.Trace.Println(chp)
			prc.SetOutChan(out)
		}
		processors = append(processors, prc)
		wg.Add(1)
		go prc.Process(in, wg)
		if idx == 0 {
			first = in
		}
	}
	return first, processors, wg, channels, hdlrs, nil
}
