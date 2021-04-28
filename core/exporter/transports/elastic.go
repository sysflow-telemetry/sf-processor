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
package transports

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	netmod "net"
	"net/http"
	"os"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	estransport "github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
)

// ElasticProto implements the TransportProtocol interface for syslog.
type ElasticProto struct {
	es     *elasticsearch.Client
	config commons.Config
}

//  NewElasticProto creates a new syslog protocol object.
func NewElasticProto(conf commons.Config) TransportProtocol {
	return &ElasticProto{config: conf}
}

// Init initializes the syslog daemon connection.
func (s *ElasticProto) Init() error {
	var err error
	cfg := elasticsearch.Config{
		Addresses: s.config.ESAddresses,
		Username:  s.config.ESUsername,
		Password:  s.config.ESPassword,
		Transport: &http.Transport{
			DialContext: (&netmod.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				//Certificates: []tls.Certificate{cert},
				//RootCAs:      caCertPool,
			},
		},
		//CACert:    ioutil.ReadFile("path/to/ca.crt"),
		Logger: &estransport.JSONLogger{Output: os.Stdout},
	}
	s.es, err = elasticsearch.NewClient(cfg)
	return err
}

// Export sends buffer to syslog daemon as an alert.
func (s *ElasticProto) Export(data commons.EncodedData) error {
	if r, ok := data.(encoders.ECSRecord); ok {
		ctx := context.Background()
		bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:         s.config.ESIndex,
			Client:        s.es,
			NumWorkers:    s.config.ESNumWorkers,   // default: 0 (= number of CPUs)
			FlushBytes:    s.config.ESFlushBuffer,  // default: 5M
			FlushInterval: s.config.ESFlushTimeout, // default: 30s
		})
		if err != nil {
			logger.Error.Println("Failed to create bulk indexer")
			return err
		}
		start := time.Now().UTC()
		body, err := json.Marshal(r)
		if err != nil {
			logger.Error.Println("Unable to marshal ECS record")
			return err
		}
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "create",
			DocumentID: r.ID,
			Body:       bytes.NewReader(body),
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				if err != nil {
					logger.Error.Print(err)
				} else {
					logger.Error.Printf("%s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		})
		if err != nil {
			logger.Error.Println("Failed to add document")
			return err
		}
		if err := bi.Close(ctx); err != nil {
			logger.Error.Printf("Failed to close bulk indexer")
			return err
		}
		duration := time.Since(start)
		biStats := bi.Stats()
		v := 1000.0 * float64(biStats.NumAdded) / float64(duration/time.Millisecond)
		logger.Info.Printf("add=%d\tflush=%d\tfail=%d\treqs=%d\tdur=%-6s\t%6d recs/s",
			biStats.NumAdded, biStats.NumFlushed, biStats.NumFailed, biStats.NumRequests,
			duration.Truncate(time.Millisecond), int64(v))
	}
	return errors.New("Expected ECSRecord as exported data")
}

// Register registers the syslog proto object with the exporter.
func (s *ElasticProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.ESTransport] = NewElasticProto
}

// Cleanup  closes the syslog connection.
func (s *ElasticProto) Cleanup() {}
