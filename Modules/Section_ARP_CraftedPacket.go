package CastHunter

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Reader for the incoming ARP RESPONSES from the requests sent out by the CRAFTER

var Network []string

func (ArpInfo *ArpInformation) CraftRead(channelmode chan struct{}) {
	source := gopacket.NewPacketSource(ArpInfo.Handler, layers.LayerTypeEthernet)
	incoming := source.Packets()
	for {
		var PKT gopacket.Packet
		select {
		case <-channelmode: // indicating a signal was received by the channel to stop
			fmt.Println("{=} closing channel")
			return
		case PKT = <-incoming:
			layer := PKT.Layer(layers.LayerTypeARP)
			if layer != nil {
				layerinf := layer.(*layers.ARP)
				if layerinf.Operation != layers.ARPReply || EqB([]byte(ArpInfo.Interface.HardwareAddr), layerinf.SourceHwAddress) {
					continue
				}
				if fmt.Sprint(net.HardwareAddr(layerinf.SourceHwAddress)) != "" && fmt.Sprint(net.IP((layerinf.SourceProtAddress))) != "" {
					MACS = append(MACS, fmt.Sprint(net.HardwareAddr(layerinf.SourceHwAddress)))
					MACS = ExterminateExtraVals(MACS)
					IPS = append(IPS, fmt.Sprint(net.IP((layerinf.SourceProtAddress))))
					IPS = ExterminateExtraVals(IPS)
				}
			}
		}
	}
}
