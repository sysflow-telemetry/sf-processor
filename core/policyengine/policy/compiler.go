package policy

type PolicyCompiler[R any] interface {
	// Compile reads one or more input policy files, parses, and translates them to internal criteria objects.
	Compile(paths ...string) ([]Rule[R], []Filter[R], error)
}
