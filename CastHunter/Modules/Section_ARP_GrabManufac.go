package CastHunter

import (
	"net"

	"github.com/google/gopacket/macs"
)

func GrabManufac(address string) (OUI string) {
	if mac, err := net.ParseMAC(address); err == nil {
		prefix := [3]byte{
			mac[0],
			mac[1],
			mac[2],
		}
		manufacturer, e := macs.ValidMACPrefixMap[prefix]
		if e {
			if manufacturer != "" {
				OUI = manufacturer
			} else {
				OUI = "Unknown"
			}
		} else {
			OUI = "Unknown"
		}
	}
	return OUI
}
