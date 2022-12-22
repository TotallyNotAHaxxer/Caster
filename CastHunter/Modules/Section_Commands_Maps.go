package CastHunter

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var GoogleHosts [][]string

var SearchQueryRoku string

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
	"roku": func() {
		if len(MACS) == len(IPS) {
			cols := []string{"IP", "MAC", "OUI"}
			var rows [][]string
			for i := 0; i < len(MACS); i++ {
				manufac := GrabManufac(MACS[i])
				if manufac == "Roku, Inc" {
					rows = append(rows, []string{
						IPS[i],
						MACS[i],
						manufac,
					})
				}
			}
			DrawTableSepColBased(rows, cols)
		} else {
			fmt.Println("[-] Calculation error: list of MACS and IPS should be the same length, this is a major issue, please report this to the github repo or developers")
		}

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
	"keyword": func(keyword string) {
		fmt.Println("[Information] Roku URL will now search for -> ", keyword, " When you choose to manipulate it")
		SearchQueryRoku = keyword
	},
}

var Enumerate = map[string]func(){
	"ipinfo": func() {
		if TargetMain != "" {
			Trace_Intrist_IPAPCO(TargetMain)
		} else {
			fmt.Println("[-] Error: Make sure the target was set using command `set target=targetIPaddress` where targetIPaddress is your targets or chrome cast devices IP")
		}
	}, // Get cast IPA Information
	"*ports": func() {
		for i := 0; i < len(IPS); i++ {
			fmt.Println("\033[34m[\033[35m*\033[34m] | Scanning -> ", IPS[i])
			GetOpenPorts(IPS[i], Range{Start: 1, End: 65535})
		}
	}, // Get all open ports of every single host
	"ports": func() {
		if TargetMain != "" {
			GetOpenPorts(TargetMain, Range{Start: 1, End: 65535})
		}
	}, // Get open ports of target
	"cast-devsaved":   func() {}, // Get cast's saved networks
	"cast-devscan":    func() {}, // Get cast to scan for wifi networks
	"cast-devforget":  func() {}, // Get cast to forget a network
	"cast-devrename":  func() {}, // Get cast to rename itself
	"cast-devreboot":  func() {}, // Get cast to reboot itself
	"cast-devfreset":  func() {}, // Get cast to factory reset itself
	"cast-devkillapp": func() {}, // Get cast to kill any application
	"cast-devsetwall": func() {}, // Get cast to set wallpaper
	"cast-devplay":    func() {}, // Get cast to play videos
	"cast-devinfo":    func() {}, // Get cast information
	"roku-devinfo": func() {
		if TargetMain != "" {
			var rows [][]string
			url := PathsRoku["devinfo"]
			newurl := fmt.Sprintf(url, TargetMain)
			resp, x := http.Get(newurl)
			if x != nil {
				fmt.Println("[+] Error when making request: -> ", x)
				return
			}
			if resp.StatusCode == 200 {
				var Structure DeviceInfo
				bv, _ := ioutil.ReadAll(resp.Body)
				xml.Unmarshal(bv, &Structure)
				rows = append(rows, []string{"udn", Structure.Udn},
					[]string{"serial-number", Structure.SerialNumber},
					[]string{"device-id", Structure.DeviceID},
					[]string{"advertising-id", Structure.AdvertisingID},
					[]string{"vendor-name", Structure.VendorName},
					[]string{"model-name", Structure.ModelName},
					[]string{"model-number", Structure.ModelNumber},
					[]string{"model-region", Structure.ModelRegion},
					[]string{"is-tv", Structure.IsTv},
					[]string{"is-stick", Structure.IsStick},
					[]string{"mobile has live tv", Structure.MobileHasLiveTv},
					[]string{"model region", Structure.ModelRegion},
					[]string{"model name", Structure.ModelName},
					[]string{"ui-resolution", Structure.UiResolution},
					[]string{"WiFi MAC", Structure.WifiMac},
					[]string{"WiFi Driver", Structure.WifiDriver},
					[]string{"has WiFi extender", Structure.HasWifiExtender},
					[]string{"has WiFi 5G support", Structure.HasWifi5GSupport},
					[]string{"has WiFi Extender", Structure.CanUseWifiExtender},
					[]string{"network-type", Structure.NetworkType},
					[]string{"signed on network", Structure.NetworkName},
					[]string{"friendly device name", Structure.FriendlyDeviceName},
					[]string{"friendly device model name", Structure.FriendlyModelName},
					[]string{"defualt device name", Structure.DefaultDeviceName},
					[]string{"user device name", Structure.UserDeviceName},
					[]string{"user device location", Structure.UserDeviceLocation},
					[]string{"build number", Structure.BuildNumber},
					[]string{"software version", Structure.SoftwareVersion},
					[]string{"software build", Structure.SoftwareBuild},
					[]string{"secure device", Structure.SecureDevice},
					[]string{"device language", Structure.Language},
					[]string{"device country ", Structure.Country},
					[]string{"device locale", Structure.Locale},
					[]string{"time zone auto", Structure.TimeZoneAuto},
					[]string{"time zone", Structure.TimeZone},
					[]string{"time zone name", Structure.TimeZoneName},
					[]string{"time zone tz", Structure.TimeZoneTz},
					[]string{"time zone offset", Structure.TimeZoneOffset},
					[]string{"clock format", Structure.ClockFormat},
					[]string{"uptime", Structure.Uptime},
					[]string{"power mode", Structure.PowerMode},
					[]string{"supports suspend", Structure.SupportsSuspend},
					[]string{"supports find remote", Structure.SupportsFindRemote},
					[]string{"supports ethernet", Structure.SupportsEthernet},
					[]string{"supports-audio-guide", Structure.SupportsAudioGuide},
					[]string{"supports-rva", Structure.SupportsRva},
					[]string{"supports-hands-free-voice-remote", Structure.HasHandsFreeVoiceRemote},
					[]string{"supports-audio-settings", Structure.SupportsAudioSettings},
					[]string{"supports-private-listening", Structure.SupportsPrivateListening},
					[]string{"supports-ecs-textedit", Structure.SupportsEcsTextedit},
					[]string{"supports-ecs-microphone", Structure.SupportsEcsMicrophone},
					[]string{"supports-wake-on-wlan", Structure.SupportsWakeOnWlan},
					[]string{"supports-airplay", Structure.SupportsAirplay},
					[]string{"supports has-play-on-roku", Structure.HasPlayOnRoku},
					[]string{"supports has-mobile-screensaver", Structure.HasMobileScreensaver},
					[]string{"trc-version", Structure.TrcVersion},
					[]string{"trc-channel", Structure.TrcChannelVersion},
					[]string{"davinci-version", Structure.DavinciVersion},
					[]string{"av-sync-calibration-enables", Structure.AvSyncCalibrationEnabled},
					[]string{"headphones-connected", Structure.HeadphonesConnected},
					[]string{"notifications-first-use", Structure.NotificationsFirstUse},
					[]string{"notifications enabled", Structure.NotificationsEnabled},
					[]string{"developer mode enabled", Structure.DeveloperEnabled})
				DrawVerticle(rows)
			}
		}

	}, // Get RoKu information
	"roku-search": func() {
		if SearchQueryRoku != "" && TargetMain != "" {
			url := PathsRoku["browse"]
			newurl := fmt.Sprintf(url, TargetMain, SearchQueryRoku)
			NewPostNoData(newurl, true)
		} else {
			fmt.Println("[-] Error: Please ensure that you have properly set a target `set target=1.1.1.1`")
			fmt.Println("[-] Error: Please ensure that you have properly set a query `set keyword=bee`")
		}
	}, // Get Roku box to go to search for a keyword
	"roku-home": func() {
		if TargetMain != "" {
			url := Keys["home"]
			newurl := fmt.Sprintf(url, TargetMain)
			NewPostNoData(newurl, true)
		} else {
			fmt.Println("[-] Error: Please ensure that you have properly set a target `set target=1.1.1.1`")
		}

	}, // Get Roku to go home
	"roku-tvinfo": func() {}, // Get Roku TV information
	"cast-*hosts": func() {}, // Enumerate all hosts that are seen as google devices (this will run every function for enumeration)
	"cast-*run":   func() {}, // Run all functions using default settings on a single cast device

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
