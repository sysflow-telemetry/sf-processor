//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
	"runtime/trace"
	"syscall"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/driver/manifest"
	"github.com/sysflow-telemetry/sf-processor/driver/pipeline"
)

var pl plugins.SFPipeline

func initSigTerm() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in terminal")
		if pl != nil {
			pl.Shutdown()
		}
	}()
}

func main() {

	// setup interruption handler
	initSigTerm()

	// setup arg parsing
	inputType := flag.String("driver", "file", fmt.Sprintf("Driver name {file|socket|<custom>}"))
	cpuprofile := flag.String("cpuprofile", "", "Write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "Write memory profile to `file`")
	traceprofile := flag.String("traceprofile", "", "Write trace profile to `file`")
	configFile := flag.String("config", "pipeline.json", "Path to pipeline configuration file")
	logLevel := flag.String("log", "info", "Log level {trace|info|warn|error}")
	driverDir := flag.String("driverdir", pipeline.DriverDir, "Dynamic driver directory")
	pluginDir := flag.String("plugdir", pipeline.PluginDir, "Dynamic plugins directory")
	version := flag.Bool("version", false, "Outputs version information")

	flag.Usage = func() {
		fmt.Println("Usage: sfprocessor [[-version]|[-driver <value>] [-log <value>] [-driverdir <value>] [-plugdir <value>] path]")
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

	// prints version information and exits
	if *version {
		hdr := sfgo.NewSFHeader()
		hdr.SetDefault(0)
		schemaVersion := hdr.Version
		fmt.Printf("Version: %s+%s, Avro Schema Version: %v, Export Schema Version: %v\n", manifest.Version, manifest.BuildNumber, schemaVersion, manifest.JSONSchemaVersion)
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
			log.Fatal("Could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *traceprofile != "" {
		f, err := os.Create(*traceprofile)
		if err != nil {
			log.Fatal("Could not create Trace profile: ", err)
			panic(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			log.Fatal("Could not create Trace profile: ", err)
			panic(err)
		}
		defer trace.Stop()
	}

	// load pipeline
	pl = pipeline.New(*driverDir, *pluginDir, *configFile)
	err := pl.Load(*inputType)
	if err != nil {
		logger.Error.Println("Unable to load pipeline error: " + err.Error())
		return
	}

	// log summary of loaded pipeline
	pl.Print()

	// initialize the pipeline
	err = pl.Init(path)
	if err != nil {
		logger.Error.Println("Error caught while starting the pipeline: " + err.Error())
		return
	}

	// memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("Could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("Could not write memory profile: ", err)
		}
	}
}
