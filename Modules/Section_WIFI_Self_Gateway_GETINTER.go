package CastHunter

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type WifiInterfaceData struct {
	InterfaceNameUsed string
	Gateway           string
}

func CheckE(x error, reason string) {
	if x != nil {
		fmt.Println(reason)
	}
}

func (Wifi *WifiInterfaceData) LoadWiFiInterface() (*net.Interface, error) {
	sysfaces, x := net.Interfaces()
	CheckE(x, "Sorry but casters WiFi module could not find a list of the interfaces...")
	var Face *net.Interface
	for _, k := range sysfaces {
		if k.Flags&net.FlagUp == 0 {
			continue
		}
		if k.Flags&net.FlagLoopback != 0 {
			continue
		}
		if k.Flags&net.FlagBroadcast == 0 {
			continue
		}
		if k.MTU == 0 {
			continue
		}
		if a, x := k.Addrs(); a == nil {
			if x != nil {
				log.Fatal(x)
			}
			continue
		}
		Face = &k
		break
	}
	if Face == nil {
		return nil, errors.New("sorry :< caster could not find any WiFi based interfaces")
	}
	Wifi.InterfaceNameUsed = Face.Name
	return Face, nil
}
