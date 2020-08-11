package driver

import (
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
)

//Driver is an interface representing telemetry drivers.
type Driver interface {
	Init(pipeline *pipeline.Pipeline) error
	Run(path string, running *bool) error
}
