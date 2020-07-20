package exporter

import (
	"strconv"
)

// Configuration keys.
const (
	ExportConfigKey      string = "export"
	ExpTypeConfigKey     string = "type"
	FormatConfigKey      string = "format"
	FlatConfigKey        string = "flat"
	ProtoConfigKey       string = "proto"
	TagConfigKey         string = "tag"
	LogSourceConfigKey   string = "source"
	HostConfigKey        string = "host"
	PortConfigKey        string = "port"
	PathConfigKey        string = "path"
	EventBufferConfigKey string = "buffer"
)

// Config defines a configuration object for the exporter.
type Config struct {
	Export      Export
	ExpType     ExportType
	Format      Format
	Flat        bool
	Proto       Proto
	Tag         string
	LogSource   string
	Host        string
	Port        int
	Path        string
	EventBuffer int
}

// CreateConfig creates a new config object from config dictionary.
func CreateConfig(conf map[string]string) Config {
	var c Config = Config{Host: "localhost", Port: 514, Path: "./export.out", Tag: "sysflow"} // default values
	if v, ok := conf[ExportConfigKey]; ok {
		c.Export = parseExportConfig(v)
	}
	if v, ok := conf[ExpTypeConfigKey]; ok {
		c.ExpType = parseExportTypeConfig(v)
	}
	if v, ok := conf[FormatConfigKey]; ok {
		c.Format = parseFormatConfig(v)
	}
	if v, ok := conf[FlatConfigKey]; ok && v == "true" {
		c.Flat = true
	}
	if v, ok := conf[ProtoConfigKey]; ok {
		c.Proto = parseProtoConfig(v)
	}
	if v, ok := conf[TagConfigKey]; ok {
		c.Tag = v
	}
	if v, ok := conf[LogSourceConfigKey]; ok {
		c.LogSource = v
	}
	if v, ok := conf[HostConfigKey]; ok {
		c.Host = v
	}
	if v, ok := conf[PortConfigKey]; ok {
		c.Port, _ = strconv.Atoi(v)
	}
	if v, ok := conf[PathConfigKey]; ok {
		c.Path = v
	}
	if v, ok := conf[EventBufferConfigKey]; ok {
		c.EventBuffer, _ = strconv.Atoi(v)
	}
	return c
}

// Export type.
type Export int

// Export config options.
const (
	StdOutExport Export = iota
	FileExport
	SyslogExport
)

func (s Export) String() string {
	return [...]string{"terminal", "file", "syslog"}[s]
}

func parseExportConfig(s string) Export {
	if StdOutExport.String() == s {
		return StdOutExport
	}
	if FileExport.String() == s {
		return FileExport
	}
	if SyslogExport.String() == s {
		return SyslogExport
	}
	return StdOutExport
}

// ExportType type.
type ExportType int

// ExportType config options.
const (
	AlertType ExportType = iota
	TelemetryType
)

func (s ExportType) String() string {
	return [...]string{"alert", "telemetry"}[s]
}

func parseExportTypeConfig(s string) ExportType {
	if AlertType.String() == s {
		return AlertType
	}
	if TelemetryType.String() == s {
		return TelemetryType
	}
	return TelemetryType
}

// Format type.
type Format int

// Format config options.
const (
	JSONFormat Format = iota
)

func (s Format) String() string {
	return [...]string{"json"}[s]
}

func parseFormatConfig(s string) Format {
	return JSONFormat
}

// Proto denotes protocol type.
type Proto int

// Proto config options.
const (
	TCPProto Proto = iota
	TCPTLSProto
	UDPProto
)

func (s Proto) String() string {
	return [...]string{"tcp", "tls", "udp"}[s]
}

func parseProtoConfig(s string) Proto {
	switch s {
	case TCPProto.String():
		return TCPProto
	case TCPTLSProto.String():
		return TCPTLSProto
	case UDPProto.String():
		return UDPProto
	}
	return TCPProto
}
