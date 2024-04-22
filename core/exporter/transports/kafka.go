//
// Copyright (C) 2024 IBM Corporation.
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

// Package transports implements transports for telemetry data.
package transports

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/commons"
)

// KafkaProto implements the TransportProtocol interface for a text file.
type KafkaProto struct {
	config   commons.Config
	producer *kafka.Producer
}

// NewKafkaProto creates a new text file protcol object.
func NewKafkaProto(conf commons.Config) TransportProtocol {
	return &KafkaProto{config: conf}
}

// Register registers the text file proto object with the exporter.
func (s *KafkaProto) Register(eps map[commons.Transport]TransportProtocolFactory) {
	eps[commons.FileTransport] = NewKafkaProto
}

// Init initializes the text file.
func (s *KafkaProto) Init() (err error) {
	s.producer, err = kafka.NewProducer(&s.config.ConfigMap)
	if err != nil {
		return fmt.Errorf("could not create kafka producer")
	}
	return
}

// Export writes the buffer to the kafka topic.
func (s *KafkaProto) Export(data []commons.EncodedData) error {
	for _, d := range data {
		if buf, ok := d.([]byte); ok {
			if err := s.producer.Produce(
				&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic:     &s.config.Topic,
						Partition: kafka.PartitionAny,
					},
					Value: buf,
				},
				nil,
			); err != nil {
				logger.Error.Printf("error producing kafka message %v", err)
			}
		}
	}
	return nil
}

// Cleanup closes the text file.
func (s *KafkaProto) Cleanup() {
	if !s.producer.IsClosed() {
		s.producer.Flush(3000)
		s.producer.Close()
	}
}
