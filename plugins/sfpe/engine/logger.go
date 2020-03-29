package engine

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// LogLevel type for enumeration.
type LogLevel int

// LogLevel enumeration.
const (
	TRACE LogLevel = iota
	INFO
	WARN
	ERROR
)

func (d LogLevel) String() string {
	return [...]string{"Trace", "Info", "Warn", "Error"}[d]
}

// Loggers reflecting different log levels.
var (
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

// InitLoggers initialize utility loggers with default i/o streams.
func InitLoggers(level LogLevel) {
	switch level {
	case TRACE:
		initLoggers(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
		break
	case INFO:
		initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
		break
	case WARN:
		initLoggers(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr)
		break
	case ERROR:
		initLoggers(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
		break
	default:
		initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}
}

func initLoggers(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warnHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		fmt.Sprintf("[%s] ", TRACE),
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		fmt.Sprintf("[%s] ", INFO),
		log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(warnHandle,
		fmt.Sprintf("[%s] ", WARN),
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		fmt.Sprintf("[%s] ", ERROR),
		log.Ldate|log.Ltime|log.Lshortfile)
}
