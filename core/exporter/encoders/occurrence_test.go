package encoders_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders"
	"github.com/sysflow-telemetry/sf-processor/core/exporter/encoders/avro/occurrence/event"
)

func TestEventSerialization(t *testing.T) {
	path := "/tmp/events.avro"
	count := 25
	fw, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.NoError(t, err)
	for i := 0; i < count; i++ {
		e := &encoders.Event{Event: event.NewEvent()}
		e.Ts = time.Now().Unix()
		e.Description = fmt.Sprintf("event %d", i)
		err := e.Serialize(fw)
		assert.NoError(t, err)
	}
	fw.Close()
	fr, err := os.OpenFile(path, os.O_RDONLY, 0644)
	assert.NoError(t, err)
	var events []*event.Event
	for {
		if e, err := event.DeserializeEvent(fr); err == nil {
			events = append(events, e)
		} else {
			break
		}
	}
	assert.Equal(t, count, len(events))
	fr.Close()
	os.Remove(path)
}
