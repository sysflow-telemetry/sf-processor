//go:build otel
// +build otel

//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Andreas Schade <san@zurich.ibm.com>
// Frederico Araujo <frederico.araujo@ibm.com>
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

// Package encoders implements codecs for exporting records and events in different data formats.
package encoders

import (
	"encoding/json"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/common"
	"google.golang.org/protobuf/proto"
)

// OtelEncoder implements an Otel encoder for telemetry records.
type OtelEncoder struct {
	config commons.Config
	batch  []commons.EncodedData
}

// NewOtelEncoder instantiates an Otel encoder.
func NewOtelEncoder(config commons.Config) Encoder {
	return &OtelEncoder{
		config: config,
		batch:  make([]commons.EncodedData, 0, config.EventBuffer)}
}

// Register registers the encoder to the codecs cache.
func (t *OtelEncoder) Register(codecs map[commons.Format]EncoderFactory) {
	codecs[commons.OtelFormat] = NewOtelEncoder
}

// Encode encodes telemetry records into an Otel representation.
func (t *OtelEncoder) Encode(r []*common.Record) ([]commons.EncodedData, error) {
	t.batch = t.batch[:0]
	for _, rec := range r {
		if otel, err := t.encode(rec); err == nil {
			t.batch = append(t.batch, otel)
		}
	}
	return t.batch, nil
}

// Encodes a telemetry record into an Otel representation.
func (t *OtelEncoder) encode(rec *common.Record) ([]byte, error) {
	var msgValue []byte
	var err error
	if t.config.Encoding == commons.JSONEncoding {
		msgValue, err = json.Marshal(rec)
		if err != nil {
			logger.Error.Println("failed to serialize record to json")
		}
	} else if t.config.Encoding == commons.ProtoEncoding {
		msgValue, err = proto.Marshal(rec)
		if err != nil {
			logger.Error.Println("failed to serialize to record proto")
		}
	} else {
		logger.Error.Printf("invalid driver encoding %s", t.config.Encoding)
	}
	return msgValue, err
}

// Cleanup cleans up resources.
func (t *OtelEncoder) Cleanup() {}
