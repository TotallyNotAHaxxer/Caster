package CastHunter

import (
	"regexp"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type SSDP_Decoded struct {
	LOCATION []string
}

var (
	LTUDP = layers.LayerTypeUDP
	LTIP  = layers.LayerTypeIPv4
)

func (Dec *SSDP_Decoded) Decoder(packet gopacket.Packet) {
	if udpl := packet.Layer(LTUDP); udpl != nil {
		udpdec, _ := udpl.(*layers.UDP)
		if udpdec.DstPort == 1900 || udpdec.SrcPort == 1900 {
			// get layer data
			// decode HTTP data
			/*
					This really is not needed in our scenario
				if il := packet.Layer(LTIP); il != nil {
				}
						if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
							if tcp := lay.(*layers.TCP); tcp != nil {
								if len(tcp.Payload) != 0 {
									r := bufio.NewReader(bytes.NewReader(tcp.Payload))
									line, x := http.ReadRequest(r)
									if x == nil {
										switch line.Proto {
										case "HTTP/1.0", "HTTP/1.1": // looking for 1.1
											if l := line.Host; l != "" {
												fmt.Println("(SSDP/HTTP[Hostname]) Found host -> ", l)
											}
											loc, x := line.Response.Location()
											if x != nil {
												return
											}
											fmt.Println("(SSDP/HTTP[Location]) Found Location -> ", loc)
										}
									}
								}
							}
						}
			*/
			// Decode payload
			if packet.ApplicationLayer().Payload() != nil {
				payload := packet.ApplicationLayer().Payload()
				// use regex search for the payloads
				regex1 := regexp.MustCompile("(?i)LOCATION:(.*)")
				res := regex1.FindAllStringSubmatch(string(payload), 1)
				for i := range res {
					newstr := strings.Trim(res[i][0], "location:LOCATION:")
					Dec.LOCATION = append(Dec.LOCATION, newstr)
				}
			}
		}
	}
}
