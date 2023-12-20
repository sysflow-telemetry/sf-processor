//
// Copyright (C) 2023 IBM Corporation.
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

// Package sysflow implements pluggable drivers for SysFlow ingestion.

package otel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"

	// "github.ibm.com/sysflow/xdr/source/xdr/generated/schema"

	v1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

const (
	fileDriverName = "otelfile"
)

func getFiles(filename string) ([]string, error) {
	var fls []string
	if fi, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	} else if fi.IsDir() {
		logger.Trace.Println("File is a directory")
		var files []fs.DirEntry
		var err error
		if files, err = os.ReadDir(filename); err != nil {
			return nil, err
		}
		for _, file := range files {
			f := filename + "/" + file.Name()
			logger.Trace.Println("File in Directory: " + f)
			fls = append(fls, f)
		}
		if len(fls) == 0 {
			return nil, errors.New("No files present in directory: " + filename)
		}

	} else {
		fls = append(fls, filename)
	}
	return fls, nil
}

// FileDriver represents reading a sysflow file from source
type FileDriver struct {
	pipeline plugins.SFPipeline
	file     *os.File
}

// NewFileDriver creates a new file driver object
func NewFileDriver() plugins.SFDriver {
	return &FileDriver{}
}

// GetName returns the driver name.
func (s *FileDriver) GetName() string {
	return fileDriverName
}

// Register registers driver to plugin cache
func (s *FileDriver) Register(pc plugins.SFPluginCache) {
	pc.AddDriver(fileDriverName, NewFileDriver)
}

// Init initializes the file driver with the pipeline
func (s *FileDriver) Init(pipeline plugins.SFPipeline, config map[string]interface{}) error {
	s.pipeline = pipeline
	return nil
}

// Run runs the file driver
func (s *FileDriver) Run(path string, running *bool) error {
	// TODO need to figure out how to call init before being able to run the GetRootChannel things
	// channel := s.pipeline.GetRootChannel()
	// otepChannel := channel.(*OTEPChannel)

	var channel OTELChannel
	otepChannel := &channel

	records := otepChannel.In

	files, err := getFiles(path)
	if err != nil {
		logger.Error.Println("Files error: ", err)
		return err
	}

	var otpLogs []*v1.ResourceLogs

	for _, fn := range files {
		logger.Trace.Println("Loading file: " + fn)
		s.file, err = os.Open(fn)
		if err != nil {
			logger.Error.Println("File open error: ", err)
			return err
		}
		bytes, err := os.ReadFile(fn)
		if err != nil {
			logger.Error.Println("File read error: ", err)
			return err
		}

		err = json.Unmarshal(bytes, &otpLogs)
		if err != nil {
			logger.Error.Println("Error unmarshaling into OTP ResourceLogs: ", err)
			return err
		}
	}

	// TODO uncomment this to figure out how to send to the parsers etc
	// for _, e := range otpLogs {
	// 	if !*running {
	// 		break
	// 	}
	// 	records <- e
	// }
	logger.Error.Println("Closing main channel")
	if records != nil {
		close(records)
	}
	return nil
}

// Cleanup tears down the driver resources.
func (s *FileDriver) Cleanup() {
	// logger.Trace.Println("Exiting ", fileDriverName)
	fmt.Println("Exiting ", fileDriverName)

	if s.file != nil {
		s.file.Close()
	}
}
