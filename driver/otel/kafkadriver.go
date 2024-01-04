package otel

import (
	"fmt"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
)

const (
	kafkaDriverName = "otelkafka"
)

type KafkaDriver struct {
	pipeline plugins.SFPipeline
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
	s.pipeline = pipeline
	return nil
}

func (s *KafkaDriver) Run(path string, running *bool) error {
	// channel := s.pipeline.GetRootChannel()
	// otelChannel, ok := channel.(*OTELChannel)
	// if !ok {
	// 	logger.Error.Println("bad root channel type")
	// 	return fmt.Errorf("bad root channel type")
	// }

	// records := otelChannel.In

	/* TODO Subscribe to Kafka here */

	return nil
}

func (s *KafkaDriver) Cleanup() {
	fmt.Println("Exiting, ", kafkaDriverName)
}
