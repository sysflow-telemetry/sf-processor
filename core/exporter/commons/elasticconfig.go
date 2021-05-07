package commons

import (
	"strconv"
	"strings"
	"time"
)

// Configuration keys.
const (
	ESAddressesConfigKey string = "es.addresses"
	ESIndexConfigKey     string = "es.index"
	ESUsernameConfigKey  string = "es.username"
	ESPasswordConfigKey  string = "es.password"
	ESWorkersConfigKey   string = "es.bulk.numWorkers"
	ESFBufferConfigKey   string = "es.bulk.flushBuffer"
	ESFTimeoutConfigKey  string = "es.bulk.flushTimeout"
)

// ESConfig holds Elastic specific configuration.
type ESConfig struct {
	ESAddresses    []string
	ESIndex        string
	ESUsername     string
	ESPassword     string
	ESNumWorkers   int
	ESFlushBuffer  int
	ESFlushTimeout time.Duration
}

// CreateElasticConfig creates a new config object from config dictionary.
func CreateElasticConfig(bc Config, conf map[string]interface{}) (c ESConfig, err error) {
	// default values
	c = ESConfig{
		ESNumWorkers:   0,
		ESFlushBuffer:  5e+6,
		ESFlushTimeout: 30 * time.Second}

	// parse config map
	if v, ok := conf[ESAddressesConfigKey].(string); ok {
		c.ESAddresses = strings.Split(v, ",")
	}
	if v, ok := conf[ESIndexConfigKey].(string); ok {
		c.ESIndex = v
	}
	if v, ok := conf[ESUsernameConfigKey].(string); ok {
		c.ESUsername = v
	} else if bc.VaultEnabled {
		s, err := bc.secrets.GetDecoded(ESUsernameConfigKey)
		if err != nil {
			return c, err
		}
		c.ESUsername = string(s)
	}
	if v, ok := conf[ESPasswordConfigKey].(string); ok {
		c.ESPassword = v
	} else if bc.VaultEnabled {
		s, err := bc.secrets.GetDecoded(ESPasswordConfigKey)
		if err != nil {
			return c, err
		}
		c.ESPassword = string(s)
	}
	if v, ok := conf[ESWorkersConfigKey].(string); ok {
		c.ESNumWorkers, _ = strconv.Atoi(v)
	}
	if v, ok := conf[ESFBufferConfigKey].(string); ok {
		c.ESFlushBuffer, _ = strconv.Atoi(v)
	}
	if v, ok := conf[ESFTimeoutConfigKey].(string); ok {
		c.ESFlushTimeout, _ = time.ParseDuration(v)
	}
	return
}
