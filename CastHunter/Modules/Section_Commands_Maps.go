package CastHunter

import "fmt"

var GoogleHosts [][]string

var View = map[string]func(){
	"hosts": func() {
		for i := 0; i < len(MACS); i++ {
			if len(MACS) == len(IPS) {
				// continue to write the table
				manufac := GrabManufac(MACS[i])
				OutputRows_HostsANDmacs = append(OutputRows_HostsANDmacs, []string{
					IPS[i],
					MACS[i],
					manufac,
				})
			} else {
				fmt.Println("Whoops there was an error, this is an issue. If you are getting this message this is a developer error, the array of MAC's FOUND should be the same length of IPS found, but it is not")
			}
		}
		cols := []string{"IP Address", "MAC Address", "Manufac"}
		if OutputRows_HostsANDmacs != nil {
			DrawTableSepColBased(OutputRows_HostsANDmacs, cols)
		} else {
			fmt.Println("[-] Sorry there are no hosts that have been discovered yet, make sure the arp module is running by outputting command `check arp`")
		}
	},
	"casts": func() {
		for i := 0; i < len(MACS); i++ {
			if len(MACS) == len(IPS) {
				// continue to write the table
				manufac := GrabManufac(MACS[i])
				if manufac == "Google, Inc." {
					GoogleHosts = append(GoogleHosts, []string{
						IPS[i],
						MACS[i],
						"Google Incorporated",
					})
				}
			} else {
				fmt.Println("Whoops there was an error, this is an issue. If you are getting this message this is a developer error, the array of MAC's FOUND should be the same length of IPS found, but it is not")
			}
		}
		cols := []string{"IP Address", "MAC", "OUI"}
		if GoogleHosts != nil {
			DrawTableSepColBased(GoogleHosts, cols)
		} else {
			fmt.Println("[-] Sorry there are no hosts that have been discovered yet, make sure the arp module is running by outputting command `check arp`")
		}
	},
}

var Check = map[string]func(){
	"arp": func() {
		fmt.Println("Arp running -> ", ArpActive)
	},
}

var Set = map[string]func(string){
	"target": func(target string) {
		fmt.Println("Set enumeration target to -> ", target)
		TargetMain = target
	},
}

var Enumerate = map[string]func(){
	"devinfo":    func() {}, // Get cast information
	"ipinfo":     func() {}, // Get cast IPA Information
	"devsaved":   func() {}, // Get cast's saved networks
	"devscan":    func() {}, // Get cast to scan for wifi networks
	"devforget":  func() {}, // Get cast to forget a network
	"devrename":  func() {}, // Get cast to rename itself
	"devreboot":  func() {}, // Get cast to reboot itself
	"devfreset":  func() {}, // Get cast to factory reset itself
	"devkillapp": func() {}, // Get cast to kill any application
	"devsetwall": func() {}, // Get cast to set wallpaper
	"devplay":    func() {}, // Get cast to play videos
}
