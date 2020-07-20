package exporter

// Event defines an interface for exported event objects.
type Event interface {
	ToJSON() []byte
	ToJSONStr() string
}
