package exporter

import (
	"crypto/tls"
	"fmt"
	syslog "github.com/RackSec/srslog"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"os"
)

const (
	SYSLOG = "syslog"
	FILE   = "file"
	TERM   = "terminal"
	NULL   = "null"
)

// ExportProtocol is an interface to support a transport protocol
type ExportProtocol interface {
	Init(conf map[string]string) error
	Export(buf []byte) error
	Register(e *Exporter)
	Cleanup()
}

// NullProto implements the ExportProtocol interface with not output
// for performance testing
type NullProto struct {
}

// NewNullProto creates a new null protocol object
func NewNullProto() ExportProtocol {
	return &NullProto{}
}

// Init intializes a new null protocol object
func (s *NullProto) Init(conf map[string]string) error {
	return nil
}

// Export does nothing
func (s *NullProto) Export(buf []byte) error {
	return nil
}

// Register registers the null protocol object with the exporter
func (s *NullProto) Register(e *Exporter) {
	e.AddExportProtocol(NULL, NewNullProto)
}

//Cleanup cleans up the null protocol object
func (s *NullProto) Cleanup() {
}

// SyslogProto implements the ExportProtocol interface for syslog
type SyslogProto struct {
	sysl   *syslog.Writer
	config Config
}

//  NewSyslogProto creates a new syslog protocol object
func NewSyslogProto() ExportProtocol {
	return &SyslogProto{}
}

// Init initializes the syslog daemon connection
func (s *SyslogProto) Init(conf map[string]string) error {
	s.config = CreateConfig(conf)
	raddr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	var err error
	if s.config.Proto == TCPTLSProto {
		// TODO: verify connection with given trust certifications
		nopTLSConfig := &tls.Config{InsecureSkipVerify: true}
		s.sysl, err = syslog.DialWithTLSConfig("tcp+tls", raddr, syslog.LOG_ALERT|syslog.LOG_DAEMON, s.config.Tag, nopTLSConfig)
	} else {
		s.sysl, err = syslog.Dial(s.config.Proto.String(), raddr, syslog.LOG_ALERT|syslog.LOG_DAEMON, s.config.Tag)
	}
	if err == nil {
		s.sysl.SetFormatter(syslog.RFC5424Formatter)
		if s.config.LogSource != sfgo.Zeros.String {
			s.sysl.SetHostname(s.config.LogSource)
		}
	}
	return err
}

// Export sends buffer to syslog daemon as an alert.
func (s *SyslogProto) Export(buf []byte) error {
	err := s.sysl.Alert(UnsafeBytesToString(buf))
	return err
}

// Register registers the syslog proto object with the exporter
func (s *SyslogProto) Register(e *Exporter) {
	e.AddExportProtocol(SYSLOG, NewSyslogProto)
}

// Cleanup  closes the syslog connection.
func (s *SyslogProto) Cleanup() {
	if s.sysl != nil {
		s.sysl.Close()
	}
}

// TextFileProto implements the ExportProtocol interface for a text file.
type TextFileProto struct {
	config  Config
	fhandle *os.File
}

//  NewTextFileProto creates a new text file protcol object
func NewTextFileProto() ExportProtocol {
	return &TextFileProto{}
}

// Init initializes the text file.
func (s *TextFileProto) Init(conf map[string]string) error {
	s.config = CreateConfig(conf)
	os.Remove(s.config.Path)
	f, err := os.OpenFile(s.config.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	s.fhandle = f
	return err
}

// Export writes the buffer to the open file.
func (s *TextFileProto) Export(buf []byte) error {
	_, err := s.fhandle.Write(buf)
	s.fhandle.WriteString("\n")
	return err
}

// Register registers the text file proto object with the exporter
func (s *TextFileProto) Register(e *Exporter) {
	e.AddExportProtocol(FILE, NewTextFileProto)
}

// Cleanup closes the text file.
func (s *TextFileProto) Cleanup() {
	if s.fhandle != nil {
		s.fhandle.Close()
	}
}

// TerminalProto implements the ExportProtocol interface of a terminal output.
type TerminalProto struct {
}

//  NewTerminalProto creates a new terminal protcol object
func NewTerminalProto() ExportProtocol {
	return &TerminalProto{}
}

//Init initializes the terminal output object
func (s *TerminalProto) Init(conf map[string]string) error {
	return nil
}

// Export exports the contets of buffer for the terminal.
func (s *TerminalProto) Export(buf []byte) error {
	fmt.Println(UnsafeBytesToString(buf))
	return nil
}

// Register registers the terminal proto object with the exporter
func (s *TerminalProto) Register(e *Exporter) {
	e.AddExportProtocol(TERM, NewTerminalProto)
}

// Cleanup cleans up the terminal output object.
func (s *TerminalProto) Cleanup() {}
