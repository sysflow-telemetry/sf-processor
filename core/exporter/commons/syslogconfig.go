package commons

import (
	"strconv"
)

// Configuration keys.
const (
	ProtoConfigKey     string = "syslog.proto"
	TagConfigKey       string = "syslog.tag"
	LogSourceConfigKey string = "syslog.source"
	HostConfigKey      string = "syslog.host"
	PortConfigKey      string = "syslog.port"
)

// SyslogConfig holds rsyslog specific configuration.
type SyslogConfig struct {
	Proto     Proto
	Tag       string
	LogSource string
	Host      string
	Port      int
}

// CreateSyslogConfig creates a new config object from config dictionary.
func CreateSyslogConfig(bc Config, conf map[string]interface{}) (c SyslogConfig, err error) {
	// default values
	c = SyslogConfig{
		Host: "localhost",
		Port: 514,
		Tag:  "sysflow"}

	// parse config map
	if v, ok := conf[ProtoConfigKey].(string); ok {
		c.Proto = parseProtoConfig(v)
	}
	if v, ok := conf[TagConfigKey].(string); ok {
		c.Tag = v
	}
	if v, ok := conf[LogSourceConfigKey].(string); ok {
		c.LogSource = v
	}
	if v, ok := conf[HostConfigKey].(string); ok {
		c.Host = v
	}
	if v, ok := conf[PortConfigKey].(string); ok {
		c.Port, _ = strconv.Atoi(v)
	}
	return
}
