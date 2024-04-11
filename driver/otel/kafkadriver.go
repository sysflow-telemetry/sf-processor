package otel

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
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
	pipeline plugins.SFPipeline
	consumer *kafka.Consumer
	encoding string
}

func NewKafkaDriver() plugins.SFDriver {
	return &KafkaDriver{}
}

func NewOtelChan(size int) interface{} {
	otc := OTELChannel{In: make(chan *otel.ResourceLogs, size)}
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
	brokerString, ok := config["otelKafkaBrokerList"]
	if !ok {
		return fmt.Errorf("no broker list found to initialize driver")
	}

	/* assumes otelKafkaTopics is a list of strings */
	topicsRaw := config["otelKafkaTopics"]
	topicsStr, ok := topicsRaw.(string)
	if !ok {
		return fmt.Errorf("invalid otelKafkaTopics list")
	}

	topicsList := strings.Split(topicsStr, ",")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokerString,
		"group.id":          "sfprocessor-otel-kafka-driver",
		"auto.offset.reset": "earliest", // might want to make this earliest
	})

	if err != nil {
		return fmt.Errorf("invalid driver config -- could not make consumer")
	}

	s.consumer = consumer
	err = s.consumer.SubscribeTopics(topicsList, nil)
	if err != nil {
		return fmt.Errorf("unable to subscribe to topics")
	}

	// set the encoding of the events read off of kafka
	enc, ok := config["encoding"]
	if !ok {
		return fmt.Errorf("invalid config -- no encoding")
	}
	encStr, ok := enc.(string)
	if (enc != "proto" && enc != "json") || !ok {
		return fmt.Errorf("invalid config -- (%s) encoding not supported", enc)
	}
	s.encoding = encStr

	s.pipeline = pipeline
	return nil
}

func (s *KafkaDriver) Run(path string, running *bool) error {
	channel := s.pipeline.GetRootChannel()
	otelChannel, ok := channel.(*OTELChannel)
	if !ok {
		logger.Error.Println("bad root channel type")
		return fmt.Errorf("bad root channel type")
	}

	records := otelChannel.In
	defer close(records)
	// defer s.pipeline.Wait()
	for {
		/* reads the message from the topics */
		msg, err := s.consumer.ReadMessage(-1)
		if err != nil {
			return fmt.Errorf("error reading message %s", err)
		}

		/* take the
		/* parses the message into an otel record log */
		dl := new(otp.LogsData)

		if s.encoding == "json" {
			err = json.Unmarshal(msg.Value, &dl)
		} else if s.encoding == "proto" {
			err = proto.Unmarshal(msg.Value, dl)
		} else {
			err = fmt.Errorf("invalid driver encoding %s", s.encoding)
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
