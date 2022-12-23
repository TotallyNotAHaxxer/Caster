package CastHunter

import "bytes"

func EqB(a, b []byte) bool {
	return bytes.Equal(a, b)
}
