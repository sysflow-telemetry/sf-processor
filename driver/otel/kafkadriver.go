package otel

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	otp "go.opentelemetry.io/proto/otlp/logs/v1"
)

const (
	kafkaDriverName = "otelkafka"
)

type KafkaDriver struct {
	pipeline plugins.SFPipeline
	consumer *kafka.Consumer
}

func NewKafkaDriver() plugins.SFDriver {
	return &KafkaDriver{}
}

func (s *KafkaDriver) GetName() string {
	return kafkaDriverName
}

func (s *KafkaDriver) Register(pc plugins.SFPluginCache) {
	pc.AddDriver(kafkaDriverName, NewKafkaDriver)
}

func (s *KafkaDriver) Init(pipeline plugins.SFPipeline, config map[string]interface{}) error {
	// TODO initalize the Kafka configuration here
	rawBL, ok := config["otelKafkaBrokerList"]
	if !ok {
		return fmt.Errorf("no broker list found to initialize driver")
	}
	/* broker list of the form { hostname: "", port: ""} */
	brokerList, ok := rawBL.([]map[string]string)
	if !ok {
		return fmt.Errorf("invalid broker list cannot initialize driver")
	}
	brokerStrs := []string{}
	for _, broker := range brokerList {
		hostname, ok := broker["hostname"]
		if !ok {
			return fmt.Errorf("invalid otelKafkaBrokerList -- no field hostname")
		}
		port, ok := broker["port"]
		if !ok {
			return fmt.Errorf("invalid otelKafkaBrokerList -- no field port")
		}
		brokerString := fmt.Sprintf("%s:%s", hostname, port)
		brokerStrs = append(brokerStrs, brokerString)
	}
	brokerString := strings.Join(brokerStrs, "-")

	/* assumes otelKafkaTopics is a list of strings */
	topicsRaw := config["otelKafkaTopics"]
	topicsStrs, ok := topicsRaw.([]string)
	if !ok {
		return fmt.Errorf("invalid otelKafkaTopics list")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokerString,
		"group.id":          "sfprocessor-otel-kafka-driver",
		"auto.offset.reset": "latest", // might want to make this earliest
	})

	if err != nil {
		return fmt.Errorf("invalid driver config -- could not make consumer")
	}

	s.consumer = consumer
	s.consumer.SubscribeTopics(topicsStrs, nil)

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

		/* parses the message into an otel record log */
		var rl *otp.ResourceLogs
		//might have to do more parsing here depends on how
		//kafka events are formed
		err = json.Unmarshal(msg.Value, rl)
		if err != nil {
			return fmt.Errorf("could not parse message")
		}
		/* sends the record to the records channel */
		records <- rl
	}
}

func (s *KafkaDriver) Cleanup() {
	fmt.Println("Exiting, ", kafkaDriverName)
}
