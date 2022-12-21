package CastHunter

import (
	"fmt"
)

var GoogleHosts [][]string

var View = map[string]func(){
	"cls": func() {
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	},
	"clear": func() {
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	},
	"artc": func() {
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
		Banner()
	},
	"scannresults": func() {
		var results Scan
		fmt.Println("| Results for - ", results.IP, " ( ", results.Hostname, " ) ")
		Cols := []string{"Ports Open"}
		var Rows [][]string
		for _, v := range results.Results {
			if v.State {
				Rows = append(Rows, []string{fmt.Sprint(v.Port)})
			}
		}
		DrawTableSepColBased(Rows, Cols)
	},
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
		fmt.Println("\033[34m[\033[35m*\033[34m] | ARP RUNNING ( ", ArpActive, " ) ")
	},
}

var Set = map[string]func(string){
	"target": func(target string) {
		fmt.Println("\033[34m[\033[35m*\033[34m] | -> ", target)
		TargetMain = target
	},
}

var Enumerate = map[string]func(){
	"devinfo": func() {}, // Get cast information
	"ipinfo": func() {
		if TargetMain != "" {
			Trace_Intrist_IPAPCO(TargetMain)
		} else {
			fmt.Println("[-] Error: Make sure the target was set using command `set target=targetIPaddress` where targetIPaddress is your targets or chrome cast devices IP")
		}
	}, // Get cast IPA Information
	"devsaved":   func() {}, // Get cast's saved networks
	"devscan":    func() {}, // Get cast to scan for wifi networks
	"devforget":  func() {}, // Get cast to forget a network
	"devrename":  func() {}, // Get cast to rename itself
	"devreboot":  func() {}, // Get cast to reboot itself
	"devfreset":  func() {}, // Get cast to factory reset itself
	"devkillapp": func() {}, // Get cast to kill any application
	"devsetwall": func() {}, // Get cast to set wallpaper
	"devplay":    func() {}, // Get cast to play videos
	"*ports": func() {
		for i := 0; i < len(IPS); i++ {
			fmt.Println("\033[34m[\033[35m*\033[34m] | Scanning -> ", IPS[i])
			GetOpenPorts(IPS[i], Range{Start: 1, End: 65535})
		}
	}, // Get all open ports of every single host
	"*hosts": func() {}, // Enumerate all hosts that are seen as google devices (this will run every function for enumeration)
	"*run":   func() {}, // Run all functions using default settings on a single cast device
	"ports": func() {
		if TargetMain != "" {
			GetOpenPorts(TargetMain, Range{Start: 1, End: 65535})
		}
	}, // Get open ports of target

}

/*
var Open []string
		if TargetMain != "" {
			timeout := time.Second * 1
			sp := 1
			ep := 65535
			var WaitGroup sync.WaitGroup
			for k := sp; k <= ep; k++ {
				WaitGroup.Add(1)
				defer WaitGroup.Done()
				connect, x := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", TargetMain, k), timeout)
				if x == nil {
					Open = append(Open, fmt.Sprint(k))
					connect.Close()
					fmt.Println("OPEN - ", k)
				}
			}
			WaitGroup.Wait()
			fmt.Println("Ports open for host - ", TargetMain)
			fmt.Println(Open)
		} else {
			fmt.Println("[-] Error: Please use the set module to set a host to scan (set target=1.1.1.1)")
		}
*/
