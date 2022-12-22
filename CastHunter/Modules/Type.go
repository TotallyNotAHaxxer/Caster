package CastHunter

import (
	"net"

	"github.com/google/gopacket/pcap"
)

type ArpInformation struct {
	Handler   *pcap.Handle
	Interface *net.Interface
}
