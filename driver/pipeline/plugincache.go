package pipeline

import (
	"errors"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/spf13/viper"
	"github.com/sysflow-telemetry/sf-apis/go/handlers"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/exporter"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
	"github.ibm.com/sysflow/sf-processor/core/processor"
	"github.ibm.com/sysflow/sf-processor/core/sfpe"
)

// PluginCache defines a data strucure for managing plugins.
type PluginCache struct {
	chanMap     map[string]interface{}
	pluginMap   map[string]*plugin.Plugin
	procFuncMap map[string]interface{}
	hdlFuncMap  map[string]interface{}
	chanFuncMap map[string]interface{}
	config      *viper.Viper
	configFile  string
}

// NewPluginCache creates a new PluginCache instance.
func NewPluginCache(conf string) *PluginCache {
	plug := &PluginCache{config: viper.New(), chanMap: make(map[string]interface{}), pluginMap: make(map[string]*plugin.Plugin), configFile: conf}
	plug.procFuncMap = map[string]interface{}{"sysflowproc": processor.NewSysFlowProc, "policyengine": sfpe.NewPolicyEngine, "exporter": exporter.NewExporter}
	plug.hdlFuncMap = map[string]interface{}{"flattener": flattener.NewFlattener}
	plug.chanFuncMap = map[string]interface{}{"sysflowchan": processor.NewSysFlowChan, "flattenerchan": flattener.NewFlattenerChan, "eventchan": sfpe.NewEventChan}
	return plug
}

// GetConfig reads the PluginCache configuration.
func (p *PluginCache) GetConfig() (*Config, error) {
	s, err := os.Stat(p.configFile)
	if os.IsNotExist(err) {
		return nil, err
	}
	if s.IsDir() {
		return nil, errors.New("pipeline config file is not a file")
	}
	dir := filepath.Dir(p.configFile)
	p.config.SetConfigName(strings.TrimSuffix(filepath.Base(p.configFile), filepath.Ext(p.configFile)))
	p.config.SetConfigType("json")
	p.config.AddConfigPath(dir)

	conf := new(Config)
	err = p.config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = p.config.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	p.updateConfigFromEnv(conf)
	return conf, nil
}

// updateConfigFromEnv updates config object with environment variables if set.
// It assumes the following convention:
// - Environment variables follow the naming schema <PROCESSOR NAME>_<CONFIG ATTRIBUTE NAME>
// - Processor name in pipeline.json is all lower case
func (p *PluginCache) updateConfigFromEnv(config *Config) {
	for _, c := range config.Pipeline {
		if proc, ok := c["processor"]; ok {
			for k, v := range p.getEnv(proc) {
				c[k] = v
			}
		}
	}
}

// getEnv returns the environemnt config settings for processor proc.
func (p *PluginCache) getEnv(proc string) map[string]string {
	var conf = make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := strings.SplitN(strings.ToLower(pair[0]), "_", 2)
		if len(key) == 2 && key[0] == proc {
			conf[key[1]] = pair[1]
		}
	}
	return conf
}

// GetPlugin retrieves a cached plugin by its name.
func (p *PluginCache) GetPlugin(mod string) (*plugin.Plugin, error) {
	var plug *plugin.Plugin
	var err error
	if val, ok := p.pluginMap[mod]; ok {
		plug = val
	} else {
		plug, err = plugin.Open(mod)
		if err != nil {
			return nil, err
		}
		p.pluginMap[mod] = plug
	}
	return plug, nil
}

// GetHandler retrieves a cached plugin handler by name.
func (p *PluginCache) GetHandler(mod string, name string) (handlers.SFHandler, error) {
	var hdl handlers.SFHandler
	if val, ok := p.hdlFuncMap[name]; ok {
		funct := val.(func() handlers.SFHandler)
		hdl = funct()
	} else {
		fName := "New" + name
		plug, err := p.GetPlugin(mod)
		if err != nil {
			return nil, err
		}
		symFlattener, err := plug.Lookup(fName)
		if err != nil {
			return nil, err
		}
		funct, ok := symFlattener.(func() handlers.SFHandler)
		if !ok {
			return nil, errors.New("Unexpected type from module symbol for handler function: " + fName)
		}
		hdl = funct()
	}
	return hdl, nil
}

// GetChan retrieves a cached plugin channel by name.
func (p *PluginCache) GetChan(mod string, ch string, size int) (interface{}, error) {
	fields := strings.Fields(ch)
	if len(fields) != 2 {
		return nil, errors.New("Channel must be of the form <identifier> <type>")
	}
	if val, ok := p.chanMap[fields[0]]; ok {
		logger.Trace.Println("Found existing channel ", fields[0])
		return val, nil
	}
	var c interface{}
	if val, ok := p.chanFuncMap[fields[1]]; ok {
		funct := val.(func(int) interface{})
		c = funct(size)
	} else {
		plug, err := p.GetPlugin(mod)
		if err != nil {
			return nil, err
		}
		fName := "New" + fields[1]
		symChan, err := plug.Lookup(fName)
		if err != nil {
			return nil, err
		}
		funct, ok := symChan.(func(int) interface{})
		if !ok {
			return nil, errors.New("Unexpected type from module symbol for channel function: " + fName)
		}
		c = funct(size)
	}
	p.chanMap[fields[0]] = c
	return c, nil
}

// GetProcessor retrieves a cached plugin processor by name.
func (p *PluginCache) GetProcessor(mod string, name string, hdl handlers.SFHandler, hdlr bool) (sp.SFProcessor, error) {
	var prc sp.SFProcessor
	if val, ok := p.procFuncMap[name]; ok {
		logger.Trace.Println("Found processor in function map: ", name)
		if hdlr {
			funct := val.(func(handlers.SFHandler) sp.SFProcessor)
			prc = funct(hdl)
		} else {
			funct := val.(func() sp.SFProcessor)
			prc = funct()
		}
	} else {
		fName := "New" + name
		plug, err := p.GetPlugin(mod)
		if err != nil {
			return nil, err
		}
		logger.Trace.Println("Plugin: ", plug)
		symProcessor, err := plug.Lookup(fName)
		if err != nil {
			return nil, err
		}
		if hdlr {
			funct, ok := symProcessor.(func(handlers.SFHandler) sp.SFProcessor)
			if !ok {
				return nil, errors.New("Unexpected type from module symbol for processor: " + fName)
			}
			prc = funct(hdl)
		} else {
			funct, ok := symProcessor.(func() sp.SFProcessor)
			if !ok {
				return nil, errors.New("Unexpected type from module symbol for processor: " + fName)
			}
			prc = funct()
		}
	}
	return prc, nil
}
