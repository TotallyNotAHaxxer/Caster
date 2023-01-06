package CastHunter

import (
	"net"
	"time"
)

// Verify a TCP port connection is possible on a given host
func VerifyPortAddress(host, port string) (ConnPossible bool) {
	h := net.JoinHostPort(host, port)
	_, x := net.DialTimeout("tcp", h, time.Duration(5*time.Second))
	if x != nil {
		ConnPossible = false
	} else {
		ConnPossible = true
	}
	return ConnPossible
}
