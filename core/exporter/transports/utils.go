package transports

import (
	"reflect"
	"unsafe"
)

// UnsafeBytesToString creates a string based on a bite slice without copying.
func unsafeBytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
