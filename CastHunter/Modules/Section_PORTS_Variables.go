package CastHunter

import "net"

type P_Results struct {
	Port  int  // Port number
	State bool // Port status or state, live or dead
}

type Range struct {
	Start int // 1
	End   int // 65535
}

type Scan struct {
	Hostname string      // Hostname
	IP       []net.IP    // IP of host
	Results  []P_Results // Scan results
}
