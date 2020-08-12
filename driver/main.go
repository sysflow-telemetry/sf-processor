//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"

	"github.ibm.com/sysflow/sf-processor/driver/sysflow"
	"github.ibm.com/sysflow/sf-processor/driver/windows"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/driver/driver"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
)

// Driver constants
const (
	SockFile   = "/var/run/sysflow.sock"
	BuffSize   = 16384
	OOBuffSize = 1024
	PluginDir  = "../resources/plugins"
)

type inputType int

const (
	file inputType = iota
	socket
	winlog
)

func (it inputType) String() string {
	return [...]string{"file", "socket", "winlog"}[it]
}

func initSigTerm(running *bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		*running = false
	}()
}

func main() {
	running := true
	initSigTerm(&running)
	// setup arg parsing
	inputType := flag.String("input", file.String(), fmt.Sprintf("Input type {%s|%s|%s}", file, socket, winlog))
	cpuprofile := flag.String("cpuprofile", "", "Write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "Write memory profile to `file`")
	configFile := flag.String("config", "pipeline.json", "Path to pipeline configuration file")
	logLevel := flag.String("log", "info", "Log level {trace|info|warn|error}")
	pluginDir := flag.String("plugdir", PluginDir, "Dynamic plugins directory")
	version := flag.Bool("version", false, "Outputs version information")

	flag.Usage = func() {
		fmt.Println("Usage: sfprocessor [[-version]|[-input <value>] [-log <value>] [-plugdir <value>] path]")
		fmt.Println()
		fmt.Println("Positional arguments:")
		fmt.Println("  path string\n\tInput path")
		fmt.Println()
		fmt.Println("Arguments:")
		flag.PrintDefaults()
		fmt.Println()
	}

	// parse args and validade positional args
	flag.Parse()
	if !*version && flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *version {
		hdr := sfgo.NewSFHeader()
		hdr.SetDefault(0)
		schemaVersion := hdr.Version
		fmt.Printf("Version: %s+%s, Avro Schema Version: %v, Export Schema Version: %v\n", Version, BuildNumber, schemaVersion, JSONSchemaVersion)
		os.Exit(0)
	}

	// retrieve positional args
	path := flag.Arg(0)

	// initialize logger
	logger.InitLoggers(logger.GetLogLevelFromValue(*logLevel))

	// CPU profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	pl := pipeline.New(*pluginDir, *configFile)
	var drv driver.Driver
	// process input
	switch *inputType {
	case file.String():
		drv = sysflow.NewFileDriver()
		break
	case socket.String():
		drv = sysflow.NewStreamingDriver()
		break
	case winlog.String():
		drv = windows.NewWinEvtDriver()
		break
	default:
		logger.Error.Println("Unrecognized input type: ", *inputType)
		os.Exit(1)
	}
	err := drv.Init(pl)
	if err != nil {
		logger.Error.Println("Driver initialization error: " + err.Error())
		return
	}

	err = pl.Load()
	if err != nil {
		logger.Error.Println("Unable to load pipeline error: " + err.Error())
		return
	}
	pl.PrintPipeline()
	err = drv.Run(path, &running)
	if err != nil {
		logger.Error.Println("Driver initialization error: " + err.Error())
		return
	}

	// memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
