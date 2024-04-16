package processor

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/otel"
	"google.golang.org/protobuf/proto"
)

const (
	pluginLabel string = "otelexporter"
)

var Plugin OTELExporter

type OTELExporter struct {
	producer    *kafka.Producer
	exportTopic string
	encoding    string
}

func NewOTELExporter() plugins.SFProcessor {
	return new(OTELExporter)
}

func (s *OTELExporter) GetName() string {
	return pluginLabel
}

func (s *OTELExporter) Init(conf map[string]interface{}) error {
	brokerString, ok := conf["otelExportKafkaBrokerList"]
	if !ok {
		return fmt.Errorf("no broker list found to initialize driver")
	}

	topicRaw, ok := conf["otelExportTopic"]
	if !ok {
		return fmt.Errorf("no topic to export to")
	}
	topicStr, ok := topicRaw.(string)
	if !ok {
		return fmt.Errorf("invalid export topic")
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerString,
		"client.id":         "sfprocessor-otel-kafka-exporter",
		"acks":              "all",
	})

	if err != nil {
		return fmt.Errorf("invalid exporter config -- could not make producer")
	}

	enc, ok := conf["encoding"]
	if !ok {
		return fmt.Errorf("invalid config -- no encoding")
	}
	encStr, ok := enc.(string)
	if (enc != "proto" && enc != "json") || !ok {
		return fmt.Errorf("invalid config -- (%s) encoding not supported", enc)
	}
	s.encoding = encStr

	s.producer = producer
	s.exportTopic = topicStr

	return nil
}

func (s *OTELExporter) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginLabel, NewOTELExporter)
}

func (s *OTELExporter) Process(ch []interface{}, wg *sync.WaitGroup) {
	for _, chi := range ch {
		cha := chi.(*plugins.Channel[*otel.ResourceLogs])
		record := cha.In
		defer wg.Done()

		for {
			fc, ok := <-record
			if !ok {
				logger.Trace.Println("Channel closed shutting down")
				break
			}
			// fmt.Printf("Dealing with a record--%s\n", fc)

			var msgValue []byte
			var err error
			if s.encoding == "json" {
				msgValue, err = json.Marshal(fc)
				if err != nil {
					logger.Trace.Println("failed to serialize record to json")
				}
			} else if s.encoding == "proto" {
				msgValue, err = proto.Marshal(fc)
				if err != nil {
					logger.Trace.Println("failed to serialize to record proto")
				}
			}

			err = s.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &s.exportTopic,
					Partition: kafka.PartitionAny,
				},
				Value: msgValue,
			},
				nil,
			)

			if err != nil {
				logger.Trace.Printf("OtelExporter Error producing kafka message %v", err)
			}

		}
	}
}

func (s *OTELExporter) SetOutChan(ch []interface{}) {}

func (s *OTELExporter) Cleanup() {}
