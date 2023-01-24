package source

// RecordChannel type
type RecordChannel[R any] struct {
	In chan R
}
