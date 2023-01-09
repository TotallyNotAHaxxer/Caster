package CastHunter

import (
	"fmt"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type ID struct {
	UUID []string
}

var Id ID

func Listener(interfacename string) {
	handler, x := pcap.OpenOffline("/home/totallynotahaxxer/Desktop/SSDPCAPTURE.pcapng")
	//handler, x := pcap.OpenLive(interfacename, 65535, true, pcap.BlockForever)
	if x != nil {
		ErrorsPackets["pcapOpenLive"]()
		fmt.Print(x)
		return
	}
	defer handler.Close()
	x = handler.SetBPFFilter(FiltersSSDP["std"])
	if x != nil {
		ErrorsPackets["filtere"]()
		fmt.Print(x)
		return
	}
	packetsrc := gopacket.NewPacketSource(handler, handler.LinkType())
	// decoder
	var D SSDP_Decoded
	for packet := range packetsrc.Packets() {
		D.Decoder(packet)
	}
	D.LOCATION = ExterminateExtraVals(D.LOCATION)
	for i := range D.LOCATION {
		url := strings.Replace(strings.TrimSpace(D.LOCATION[i]), "\r", "", -1)
		Id.UUID = append(Id.UUID, ExtractUUID(url, true))
	}

}
