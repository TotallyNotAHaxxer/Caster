package CastHunter

import (
	"fmt"
	"net"
	"sync"
)

func Hunter(singleinterface bool) {
	var wg sync.WaitGroup
	interfaces, x := net.Interfaces()
	if x != nil {
		fmt.Println(x)
	} else {
		if singleinterface {
			if len(interfaces) > -1 {
				// find any trashy networks or interfaces that do not work out such as lo
				for i := 0; i < len(interfaces); i++ {
					if interfaces[i].Name == "lo" {
						interfaces = append(interfaces[:i], interfaces[i+1:]...)
						i--
					}
				}
				fmt.Printf("\n\033[38;5;50m[\033[38;5;15mInformation\033[38;5;50m] Using Interface (\033[38;5;15m%v\033[38;5;50m) ", interfaces[0].Name)
				wg.Add(1)
				interfacename := interfaces[0]
				go func(interfacename net.Interface) {
					defer wg.Done()
					if x := Scanner(&interfacename); x != nil {
						fmt.Println("[-] Error Scanning on interface -> ", interfacename.Name, x)
					}

				}(interfacename)
			}
		} else {
			for _, interfacename := range interfaces {
				wg.Add(1)
				go func(interfacename net.Interface) {
					defer wg.Done()
					if x := Scanner(&interfacename); x != nil {
						fmt.Println("[-] Error Scanning on interface -> ", interfacename.Name, x)
					}

				}(interfacename)
			}
		}
	}
	wg.Wait()
}

// Grab and return interface names
func RetInterface() (interfacename net.Interface) {
	interfaces, x := net.Interfaces()
	if x != nil {
		fmt.Println(x)
	} else {
		if len(interfaces) > -1 {
			// find any trashy networks or interfaces that do not work out such as lo
			for i := 0; i < len(interfaces); i++ {
				if interfaces[i].Name == "lo" {
					interfaces = append(interfaces[:i], interfaces[i+1:]...)
					i--
				}
				interfacename = interfaces[0]
			}
		}
	}
	return interfacename
}
