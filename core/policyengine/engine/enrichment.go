package engine

// Handler defines an interface for SysFlow enrichment routines.
type Handler interface {
	Init(confPath string) error
	ProcessSync(r *Record) (interface{}, error)
	ProcessAsync(r *Record, callback func(o interface{})) error
	Cleanup() error
}
