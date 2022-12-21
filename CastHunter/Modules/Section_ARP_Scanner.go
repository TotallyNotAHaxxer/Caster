package CastHunter

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket/pcap"
)

func Scanner(i *net.Interface) (x error) {
	var functioninformation ArpInformation
	var handler *pcap.Handle
	var Network *net.IPNet
	if addresses, e := i.Addrs(); e != nil {
		return e
	} else {
		for _, netAddress := range addresses {
			if internetP, ok := netAddress.(*net.IPNet); ok {
				if ipv4 := internetP.IP.To4(); ipv4 != nil {
					Network = &net.IPNet{
						IP:   ipv4,
						Mask: internetP.Mask[len(internetP.Mask)-4:],
					}
					break
				}
			}

		}
	}
	if !CheckAddress(Network) {
		return
	}
	fmt.Printf("\n\033[38;5;50m[\033[38;5;15mInformation\033[38;5;50m] Using network range (\033[38;5;15m%v\033[38;5;50m) for interface (\033[38;5;15m%v\033[38;5;50m) ", Network, i.Name)
	handler, x = pcap.OpenLive(i.Name, 65536, true, pcap.BlockForever)
	if x != nil {
		return x
	}
	defer handler.Close()
	ender := make(chan struct{})
	functioninformation.Handler = handler
	functioninformation.Interface = i
	go functioninformation.CraftRead(ender)
	defer close(ender)
	for {
		if x = functioninformation.Crafter(Network); x != nil {
			fmt.Println("Error when writing packets -> ", functioninformation.Interface.Name, x)
			return x
		}
		time.Sleep(time.Duration(SleepInterval) * time.Second)
	}
}
