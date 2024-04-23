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

// Package otel implements pluggable drivers for otel ingestion.
package otel

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/otel"
	otp "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/protobuf/proto"
)

const (
	kafkaDriverName  = "otelkafka"
	kafkaChannelName = "otelchan"
)

type KafkaDriver struct {
	config   KafkaConfig
	pipeline plugins.SFPipeline
	consumer *kafka.Consumer
}

func NewKafkaDriver() plugins.SFDriver {
	return &KafkaDriver{}
}

func NewOtelChan(size int) interface{} {
	otc := plugins.Channel[*otel.ResourceLogs]{In: make(chan *otel.ResourceLogs, size)}
	return &otc
}

func (s *KafkaDriver) GetName() string {
	return kafkaDriverName
}

func (s *KafkaDriver) Register(pc plugins.SFPluginCache) {
	pc.AddDriver(kafkaDriverName, NewKafkaDriver)
	pc.AddChannel(kafkaChannelName, NewOtelChan)
}

func (s *KafkaDriver) Init(pipeline plugins.SFPipeline, config map[string]interface{}) error {
	conf, err := CreateKafkaConfig(config)
	if err != nil {
		return fmt.Errorf("caught error while reading kafka driver configuration")
	}

	consumer, err := kafka.NewConsumer(&conf.ConfigMap)
	if err != nil {
		return fmt.Errorf("could not create kafka consumer")
	}

	err = consumer.SubscribeTopics(conf.Topics, nil)
	if err != nil {
		return fmt.Errorf("unable to subscribe to kafka topics: %v", conf.Topics)
	}

	s.config = conf
	s.consumer = consumer
	s.pipeline = pipeline

	return nil
}

func (s *KafkaDriver) Run(path string, running *bool) error {
	channel := s.pipeline.GetRootChannel()
	otelChannel, ok := channel.(*plugins.Channel[*otel.ResourceLogs])
	if !ok {
		return fmt.Errorf("bad root channel type")
	}

	records := otelChannel.In
	defer close(records)

	for {
		/* reads the message from the topics */
		msg, err := s.consumer.ReadMessage(-1)
		if err != nil {
			return fmt.Errorf("error reading message %s", err)
		}

		/* parses the message into an otel record log */
		dl := new(otp.LogsData)

		if s.config.Encoding == JSONEncoding {
			err = json.Unmarshal(msg.Value, &dl)
		} else {
			err = proto.Unmarshal(msg.Value, dl)
		}
		if err != nil {
			return fmt.Errorf("could not parse message %v", err)
		}

		/* sends the record to the records channel */
		for _, rl := range dl.ResourceLogs {
			records <- rl
		}
	}
}

func (s *KafkaDriver) Cleanup() {
	fmt.Println("Exiting, ", kafkaDriverName)
}
