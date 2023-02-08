package source_test

import (
	"encoding/base64"
	"testing"
)

func TestBase64Offset(t *testing.T) {
	msg := "http://"
	offset0 := base64.StdEncoding.EncodeToString([]byte(msg))
	t.Logf("%s, %s", msg, offset0)

	// for i := 0; i < 3; i++ {
	// 	d := append([]byte{" "}, )
	// 	offset := base64.StdEncoding.EncodeToString()
	// }

	msg1 := []byte(msg)[1 : len(msg)-3]
	offset1 := base64.StdEncoding.EncodeToString(append([]byte(" "), msg1...))
	t.Logf("%s, %s", msg1, offset1)

	msg2 := []byte(msg)[2:]
	offset2 := base64.StdEncoding.EncodeToString(append([]byte("  "), msg2...))
	t.Logf("%s, %s", msg2, offset2)

	// msg = "ttp://"
	// offset := base64.StdEncoding.EncodeToString([]byte(msg))
	// t.Logf("%s, %s", msg, offset)

	// msg = "tp://"
	// offset = base64.StdEncoding.EncodeToString([]byte(msg))
	// t.Logf("%s, %s", msg, offset)

	// start_offsets := []int{0, 2, 3}
	// end_offsets := []int{0, -3, -2}

}
