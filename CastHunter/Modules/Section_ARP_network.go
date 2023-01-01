package CastHunter

import (
	"encoding/binary"
	"net"
)

// Arp section of CastHunter

func GrabNetwork(netnet *net.IPNet) []net.IP {
	var networkingfu []net.IP
	networknumber := binary.BigEndian.Uint32([]byte(netnet.IP))
	networkmask := binary.BigEndian.Uint32([]byte(netnet.Mask))
	networkwhole := networknumber & networkmask
	networkbroad := networkwhole | ^networkmask
	for networkwhole++; networkwhole < networkbroad; networkwhole++ {
		var buffer [4]byte // holds 4 octets
		binary.BigEndian.PutUint32(buffer[:], networkwhole)
		networkingfu = append(networkingfu, net.IP(buffer[:]))
	}
	return networkingfu
}
