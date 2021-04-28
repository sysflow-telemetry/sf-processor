package transports

import "github.com/sysflow-telemetry/sf-processor/core/exporter/commons"

type ESConfig struct {
	commons.Config
	myatt string
}

func (c *ESConfig) test() {

}
