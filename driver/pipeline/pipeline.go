package pipeline

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/logger"
)

// Pipeline represents a loaded plugin pipeline
type Pipeline struct {
	wg          *sync.WaitGroup
	processors  []plugins.SFProcessor
	channels    []interface{}
	hdlrs       []plugins.SFHandler
	pluginCache *PluginCache
	config      string
	pluginDir   string
}

// New creates a new pipeline object
func New(pluginDir string, config string) *Pipeline {
	return &Pipeline{config: config,
		pluginDir:   pluginDir,
		wg:          new(sync.WaitGroup),
		pluginCache: NewPluginCache(config),
	}

}

// GetNumChannels returns the number of channels in the pipeline
func (pl *Pipeline) GetNumChannels() int {
	return len(pl.channels)
}

// GetNumProcessors returns the number of processors in the pipeline
func (pl *Pipeline) GetNumProcessors() int {
	return len(pl.processors)
}

// GetNumHandlers returns the number of handlers in the pipeline
func (pl *Pipeline) GetNumHandlers() int {
	return len(pl.hdlrs)
}

// GetPluginCache returns the plugin cache for the pipeline
func (pl *Pipeline) GetPluginCache() *PluginCache {
	return pl.pluginCache
}

// AddChannel adds a channel to the plugin cache
func (pl *Pipeline) AddChannel(channelName string, channel interface{}) {
	pl.pluginCache.AddChannel(channelName, channel)
}

// Load loads and enables the pipeline
func (pl *Pipeline) Load() error {
	if err := pl.pluginCache.LoadPlugins(pl.pluginDir); err != nil {
		logger.Error.Println("Unable to load dynamic plugins: ", err)
		return err
	}

	conf, err := pl.pluginCache.GetConfig()
	if err != nil {
		logger.Error.Println("Unable to load pipeline config: ", err)
		return err
	}
	setManifestInfo(conf)
	var in interface{}
	var out interface{}
	for _, p := range conf.Pipeline {
		hdler := false
		var hdl plugins.SFHandler
		if val, ok := p[HdlConfig]; ok {
			hdl, err = pl.pluginCache.GetHandler(val)
			if err != nil {
				logger.Error.Println(err)
				return err
			}
			pl.hdlrs = append(pl.hdlrs, hdl)
			xType := fmt.Sprintf("%T", hdl)
			logger.Trace.Println(xType)
			hdler = true
		}
		var prc plugins.SFProcessor
		if val, ok := p[ProcConfig]; ok {
			prc, err = pl.pluginCache.GetProcessor(val, hdl, hdler)
			if err != nil {
				logger.Error.Println(err)
				return err
			}
			tp := fmt.Sprintf("%T", prc)
			logger.Trace.Println(tp)
			err = prc.Init(p)
			if err != nil {
				logger.Error.Println(err)
				return err
			}
		} else {
			logger.Error.Println("processor or handler tag must exist in plugin config")
			return err
		}
		if v, o := p[InChanConfig]; o {
			in, err = pl.pluginCache.GetChan(v, ChanSize)
			pl.channels = append(pl.channels, in)
			chp := fmt.Sprintf("%T", in)
			logger.Trace.Println(chp)
		} else {
			logger.Error.Println("in tag must exist in plugin config")
			return errors.New("in tag must exist in plugin config")
		}
		if v, o := p[OutChanConfig]; o {
			out, err = pl.pluginCache.GetChan(v, ChanSize)
			chp := fmt.Sprintf("%T", out)
			pl.channels = append(pl.channels, out)
			logger.Trace.Println(chp)
			prc.SetOutChan(out)
		}
		pl.processors = append(pl.processors, prc)
		pl.wg.Add(1)
		go prc.Process(in, pl.wg)
	}
	return nil
}

// GetRootChannel returns the first channel in the pipeline
func (pl *Pipeline) GetRootChannel() interface{} {
	if len(pl.channels) > 0 {
		return pl.channels[0]
	}
	return nil
}

func (pl *Pipeline) PrintPipeline() {
	logger.Trace.Printf("Loaded %d stages\n", len(pl.processors))
	logger.Trace.Printf("Loaded %d channels\n", len(pl.channels))
	logger.Trace.Printf("Loaded %d hdlrs\n", len(pl.hdlrs))
}

func (pl *Pipeline) Wait() {
	pl.wg.Wait()
}
