package encoders_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/linkedin/goavro"
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

func TestEventContainerSerialization(t *testing.T) {
	path := "/tmp/events_schema.avro"
	count := 25
	fw, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.NoError(t, err)
	cw, err := event.NewEventWriter(fw, container.Snappy, 512)
	assert.NoError(t, err)
	for i := 0; i < count; i++ {
		e := &encoders.Event{Event: event.NewEvent()}
		e.Ts = time.Now().Unix()
		e.Description = fmt.Sprintf("event %d", i)
		err := cw.WriteRecord(e)
		assert.NoError(t, err)
	}
	cw.Flush()
	fw.Close()
	fr, err := os.OpenFile(path, os.O_RDONLY, 0644)
	assert.NoError(t, err)
	cr, err := event.NewEventReader(fr)
	assert.NoError(t, err)
	var events []*event.Event
	for {
		if e, err := cr.Read(); err == nil {
			events = append(events, e)
		} else {
			break
		}
	}
	assert.Equal(t, count, len(events))
	fr.Close()
	os.Remove(path)
}

func TestGoavroEventSerialization(t *testing.T) {
	path := "/tmp/events_goavro.avro"
	count := 25

	fw, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	assert.NoError(t, err)

	buf, err := ioutil.ReadFile("./avro/occurrence/avsc/Event.avsc")
	assert.NoError(t, err)
	codec, err := goavro.NewCodec(string(buf))
	assert.NoError(t, err)

	ocfw, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:               fw,
		Codec:           codec,
		CompressionName: "snappy",
	})
	assert.NoError(t, err)

	var values []map[string]interface{}
	for i := 0; i < count; i++ {
		e := &encoders.Event{Event: event.NewEvent()}
		e.Ts = time.Now().Unix()
		e.Description = fmt.Sprintf("event %d", i)
		var m map[string]interface{}
		s, _ := json.Marshal(e)
		json.Unmarshal(s, &m)
		values = append(values, m)
	}
	err = ocfw.Append(values)
	assert.NoError(t, err)

	fr, err := os.OpenFile(path, os.O_RDONLY, 0644)
	assert.NoError(t, err)
	ocfr, err := goavro.NewOCFReader(fr)
	assert.NoError(t, err)
	var events []interface{}
	for ocfr.Scan() {
		d, err := ocfr.Read()
		assert.NoError(t, err)
		events = append(events, d)
	}
	assert.Equal(t, count, len(events))

	values = nil
	for i := 0; i < count; i++ {
		e := &encoders.Event{Event: event.NewEvent()}
		e.Ts = time.Now().Unix()
		e.Description = fmt.Sprintf("event (2) %d", i)
		var m map[string]interface{}
		s, _ := json.Marshal(e)
		json.Unmarshal(s, &m)
		values = append(values, m)
	}
	err = ocfw.Append(values)
	assert.NoError(t, err)

	fr, err = os.OpenFile(path, os.O_RDONLY, 0644)
	assert.NoError(t, err)
	ocfr, err = goavro.NewOCFReader(fr)
	assert.NoError(t, err)
	events = nil
	for ocfr.Scan() {
		d, err := ocfr.Read()
		assert.NoError(t, err)
		events = append(events, d)
	}
	assert.Equal(t, 2*count, len(events))

	fw.Close()
	fr.Close()
	os.Remove(path)
}
