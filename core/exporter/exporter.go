//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Andreas Schade <san@zurich.ibm.com>
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
package exporter

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
        netmod "net"
	"net/http"
	"os"
	"sync"
	"time"

	syslog "github.com/RackSec/srslog"
        elasticsearch "github.com/elastic/go-elasticsearch/v8"
        estransport "github.com/elastic/go-elasticsearch/v8/estransport"
        esutil "github.com/elastic/go-elasticsearch/v8/esutil"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/engine"
)

const (
	pluginName string = "exporter"
)

// Exporter defines a syslogger plugin.
type Exporter struct {
	recs    []*engine.Record
	counter int
	sysl    *syslog.Writer
        es	*elasticsearch.Client
	config  Config
}

// NewExporter creates a new plugin instance.
func NewExporter() plugins.SFProcessor {
	return new(Exporter)
}

// GetName returns the plugin name.
func (s *Exporter) GetName() string {
	return pluginName
}

// Register registers plugin to plugin cache.
func (s *Exporter) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewExporter)
}

// Init initializes the plugin with a configuration map and cache.
func (s *Exporter) Init(conf map[string]interface{}) error {
	var err error
	s.config = CreateConfig(conf)
	if s.config.Export == FileExport {
		os.Remove(s.config.Path)
	} else if s.config.Export == SyslogExport {
		raddr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
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
	} else if s.config.Export == ESExport {
		cfg := elasticsearch.Config{
			Addresses: s.config.ESAddresses,
			Username:  s.config.ESUsername,
			Password:  s.config.ESPassword,
			Transport: &http.Transport{
				//MaxIdleConnsPerHost:   10,
				//ResponseHeaderTimeout: time.Second,
				DialContext:           (&netmod.Dialer{Timeout: time.Second}).DialContext,
				TLSClientConfig: &tls.Config{
					//MinVersion: tls.VersionTLS11,
					InsecureSkipVerify: true,
					//Certificates: []tls.Certificate{cert},
					//RootCAs:      caCertPool,
					// ...
				},
			},
			//CACert:    ioutil.ReadFile("path/to/ca.crt"),
			Logger:    &estransport.JSONLogger{ Output: os.Stdout },
			//Logger:    &estransport.ColorLogger{ Output: os.Stdout, EnableRequestBody: true },
		}

		s.es, err = elasticsearch.NewClient(cfg)
		if err == nil {
			logger.Info.Printf("Successfully created ES client for endpoints: %v", cfg.Addresses)
		}
	}
	return err
}

// Process implements the main interface of the plugin.
func (s *Exporter) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*engine.RecordChannel)
	record := cha.In
	defer wg.Done()

	maxIdle := 1 * time.Second
	ticker := time.NewTicker(maxIdle)
	defer ticker.Stop()
	lastFlush := time.Now()

	logger.Trace.Printf("Starting Exporter in mode %s with channel capacity %d", s.config.Export.String(), cap(record))
RecLoop:
	for {
		select {
		case fc, ok := <-record:
			if ok {
				s.counter++
				s.recs = append(s.recs, fc)
				if s.counter >= s.config.EventBuffer {
					s.process()
					s.recs = s.recs[:0]
					s.counter = 0
					lastFlush = time.Now()
				}
			} else {
				s.process()
				logger.Trace.Println("Channel closed. Shutting down.")
				break RecLoop
			}
		case <-ticker.C:
			// force flush records after 1sec idle
			if time.Now().Sub(lastFlush) > maxIdle && s.counter > 0 {
				s.process()
				s.recs = s.recs[:0]
				s.counter = 0
				lastFlush = time.Now()
			}
		}
	}
}

func (s *Exporter) process() {
	s.export(s.createEvents())
}

func (s *Exporter) createEvents() []Event {
	if s.config.ExpType == BatchType {
		return CreateOffenses(s.recs, s.config)
	}
	return CreateTelemetryRecords(s.recs, s.config)
}

func (s *Exporter) export(events []Event) {
	if s.config.Format == JSONFormat || s.config.Format == ECSFormat {
		s.exportAsJSON(events)
	}
}

func (s *Exporter) exportAsJSON(events []Event) {
	switch s.config.Export {
	case StdOutExport:
		for _, evt := range events {
			fmt.Println(evt.ToJSONStr())
		}
	case SyslogExport:
		for _, evt := range events {
			if err := s.sysl.Alert(evt.ToJSONStr()); err != nil {
				logger.Error.Println("Can't export to syslog:\n", err)
				break
			}
		}
	case FileExport:
		f, err := os.OpenFile(s.config.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Error.Println("Can't open trace file:", err)
			break
		}
		defer f.Close()
		for _, evt := range events {
			if _, err := fmt.Fprintln(f, evt.ToJSONStr()); nil != err {
				logger.Error.Println("Can't write to trace file:\n", err)
				break
			}
		}
        case ESExport:
		logger.Info.Printf("Bulk size: %d events", len(events))

		ctx := context.Background()

		bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:         s.config.ESIndex,
			Client:        s.es,
			NumWorkers:    s.config.ESNumWorkers,     // default: 0 (= number of CPUs)
			FlushBytes:    s.config.ESFlushBuffer,    // default: 5M
			FlushInterval: s.config.ESFlushTimeout,   // default: 30s
		})
		if err != nil {
			logger.Error.Printf("Failed to create bulk indexer: %s", err)
                        return
		}

		start := time.Now().UTC()

		for _, evt := range events {
			err = bi.Add(ctx, esutil.BulkIndexerItem{
				Action:     "create",
				DocumentID: evt.ID(),
				Body:       bytes.NewReader(evt.ToJSON()),
				OnFailure:  func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						logger.Error.Print(err)
					} else {
						logger.Error.Printf("%s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			})
			if err != nil {
				logger.Error.Printf("Failed to add document: %s", err)
			}
		}

		if err := bi.Close(ctx); err != nil {
			logger.Error.Printf("Failed to close bulk indexer: %s", err)
			return
		}

		duration := time.Since(start)
		biStats := bi.Stats()
		v := 1000.0 * float64(biStats.NumAdded) / float64(duration/time.Millisecond)
		logger.Info.Printf("add=%d\tflush=%d\tfail=%d\treqs=%d\tdur=%-6s\t%6d recs/s",
			biStats.NumAdded, biStats.NumFlushed, biStats.NumFailed, biStats.NumRequests,
			duration.Truncate(time.Millisecond), int64(v))
	}
}

func Sha256Hex(val []byte) string {
        hash := sha256.Sum256(val)
        return hex.EncodeToString(hash[0:sha256.Size])
}

// SetOutChan sets the output channel of the plugin.
func (s *Exporter) SetOutChan(ch []interface{}) {}

// Cleanup tears down plugin resources.
func (s *Exporter) Cleanup() {
	logger.Trace.Println("Exiting ", pluginName)
}

// This function is not run when module is used as a plugin.
func main() {}
