package CastHunter

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Craft the layers and write the the request on the given network
// this is a broadcast based ARP, it writes to every device on the network

func (ArpInfo *ArpInformation) Crafter(Addr *net.IPNet) (x error) {
	LAYER_ETHERNET := layers.Ethernet{
		SrcMAC: ArpInfo.Interface.HardwareAddr,
		DstMAC: net.HardwareAddr{
			NULL,
			NULL,
			NULL,
			NULL,
			NULL,
			NULL,
		},
		EthernetType: layers.EthernetTypeARP,
	}
	LAYERS_ARP := layers.ARP{
		DstHwAddress: []byte{
			MVB,
			MVB,
			MVB,
			MVB,
			MVB,
			MVB,
		},
		SourceProtAddress: []byte(
			Addr.IP,
		),
		SourceHwAddress: []byte(
			ArpInfo.Interface.HardwareAddr,
		),
		Operation:       layers.ARPRequest,
		ProtAddressSize: 4, // STD IPA SIZE (4: 1.1.1.1)
		HwAddressSize:   6, // STD MAC SIZE (6: ff:ff:ff:ff:ff:ff)
		Protocol:        layers.EthernetTypeIPv4,
		AddrType:        layers.LinkTypeEthernet,
	}
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	for _, address := range GrabNetwork(Addr) {
		LAYERS_ARP.DstProtAddress = []byte(address)
		gopacket.SerializeLayers(buffer, options, &LAYER_ETHERNET, &LAYERS_ARP)
		if x = ArpInfo.Handler.WritePacketData(buffer.Bytes()); x != nil {
			return x
		}
	}
	return x
}
