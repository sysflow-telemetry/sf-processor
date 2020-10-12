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
package windows

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.ibm.com/sysflow/goutils/logger"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/winlogbeat/checkpoint"
	"github.com/elastic/beats/v7/winlogbeat/eventlog"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
	"github.ibm.com/sysflow/sf-processor/driver/windows/sysmon"
)

const (
	driverName  = "winlog"
	channelName = "flattenerchan"
)

const (
	cName        = "name"
	cAPI         = "api"
	cChannelSize = 100000
)

// WinEvtDriver represents a Windows Event Data Source
type WinEvtDriver struct {
	pipeline plugins.SFPipeline
}

// NewWinEvtDriver creates a new windows event driver
func NewWinEvtDriver() plugins.SFDriver {
	return &WinEvtDriver{}
}

// GetName returns the driver name.
func (s *WinEvtDriver) GetName() string {
	return driverName
}

// Register registers driver to plugin cache
func (s *WinEvtDriver) Register(pc plugins.SFPluginCache) {
	pc.AddDriver(driverName, NewWinEvtDriver)
}

// Init initializes the driver.
func (s *WinEvtDriver) Init(pipeline plugins.SFPipeline) error {
	s.pipeline = pipeline
	s.pipeline.AddChannel(channelName, flattener.NewFlattenerChan)
	return nil
}

// Run processes windows event logs and creates and exports them into the pipeline
func (s *WinEvtDriver) Run(path string, running *bool) error {
	conf := make(map[string]interface{})
	in := s.pipeline.GetRootChannel().(*flattener.FlatChannel)
	sm := sysmon.NewSMProcessor(in)
	conf[cAPI] = ""
	conf[cName] = sm.GetProvider()
	cfg, err := common.NewConfigFrom(conf)
	if err != nil {
		logger.Error.Println("winevtlog provider config error:", err)
		return err
	}

	eventLog, err := eventlog.New(cfg)
	if err != nil {
		logger.Error.Println("Failed to create new event log error:", err)
		return err
	}
	cp := checkpoint.EventLogState{}
	err = eventLog.Open(cp)
	if err != nil {
		fmt.Printf("failed to open windows event log: %v", err)
		return err
	}
	logger.Trace.Printf("Windows Event Log '%s' opened successfully", eventLog.Name())
	// setup closing the API if either the run function is signaled asynchronously
	// to shut down or when returning after io.EOF
	/*cancelCtx, cancelFn := ctxtool.WithFunc(ctxtool.FromCanceller(ctx.Cancelation), func() {
		if err := api.Close(); err != nil {
			log.Errorf("Error while closing Windows Eventlog Access: %v", err)
		}
	})
	defer cancelFn()
	*/
	// read loop
	for *running {
		records, err := eventLog.Read()
		switch err {
		case nil:
			break
		case io.EOF:
			logger.Trace.Printf("End of Winlog event stream reached: %v", err)
			return nil
		default:
			// only log error if we are not shutting down
			/*	if cancelCtx.Err() != nil {
				return nil
			}*/

			logger.Error.Printf("Error occured while reading from Windows Event Log '%v': %v", eventLog.Name(), err)
			return errors.New("Unknown Error occurred while looping through windows evt")
		}

		if len(records) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		sm.Process(records)
	}
	close(in.In)
	s.pipeline.Wait()
	return nil
}
