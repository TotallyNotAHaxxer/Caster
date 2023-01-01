package CastHunter

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var GoogleHosts [][]string

var SearchQueryRoku string
var ApplicationIDROKU string
var DevName string
var WPAID string

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
		var newrow [][]string
		for i := 0; i < len(MACS); i++ {
			if len(MACS) == len(IPS) {
				// continue to write the table
				manufac := GrabManufac(MACS[i])
				newrow = append(newrow, []string{
					IPS[i],
					MACS[i],
					manufac,
				})
			}
		}
		cols := []string{"IP Address", "MAC Address", "Manufac"}
		if newrow != nil {
			DrawTableSepColBased(newrow, cols)
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
	"server": func() {
		conn, x := net.Dial("tcp", "localhost:5429")
		if x != nil {
			ErrorHandler[100]()
			fmt.Println(x)
			fmt.Print("\n\n\n")
			return
		} else {
			var messsages string
			messsages += "Server Up?       => TRUE\n"
			messsages += "Server Port      => " + "5429" + "\n"
			messsages += "Server Host      => http://localhost:5429"
			DrawUtilsBox(messsages)
			defer conn.Close()
		}
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
	"appid": func(appid string) {
		fmt.Println("[information] The ROKU APPID will be set to -> ", appid, " when you choose to manipulate it ")
		ApplicationIDROKU = appid
	},
	"devname": func(devname string) {
		fmt.Println("[information] The GoogleCasts name will be changed to -> ", devname, " when you choose to enumerate")
		DevName = devname
	},
	"wpaid": func(wpaid string) {
		fmt.Println("[Information] The GoogleCasts WPA ID will be changed to -> ", wpaid, " when you choose to tell the cast to forget the ID")
		WPAID = wpaid
	},
	"remrok": func(remrok string) {
		RemAddr = remrok
		fmt.Println("Removing -> ", RemAddr, " from list when you choose to run the removing option")
	},
}

var Enumerate = map[string]func(){
	"test": func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				return
			}
		}()
		index := []int{1, 2, 3, 4, 5}
		fmt.Println(index[len(index)])

	},
	"ipinfo": func() {
		if TargetMain != "" {
			Trace_Intrist_IPAPCO(TargetMain)
		} else {
			ErrorHandler[80]()
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
		} else {
			ErrorHandler[80]()
		}
	}, // Get open ports of target
	"cast-devsaved": func() {
		if TargetMain != "" {
			url := PathsCast["configurednet"]
			newurl := fmt.Sprintf(url, TargetMain, DevicePorts["https"])
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := http.Client{
				Transport: tr,
			}
			resp, x := client.Get(newurl)
			if x != nil {
				ErrorHandler[50]()
				fmt.Print(x)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				type Wifi []struct {
					Ssid      string `json:"ssid"`
					WpaAuth   int    `json:"wpa_auth"`
					WpaCipher int    `json:"wpa_cipher"`
					WpaID     int    `json:"wpa_id"`
				}
				var b Wifi
				bv, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(bv, &b)
				var rows [][]string
				cols := []string{"SSID", "WPA_AUTH", "WPA_CIPHER", "WPAID"}
				for i := 0; i < len(b); i++ {
					rows = append(rows,
						[]string{b[i].Ssid},
						[]string{fmt.Sprint(b[i].WpaAuth)},
						[]string{fmt.Sprint(b[i].WpaCipher)},
						[]string{fmt.Sprint(b[i].WpaID)},
					)
				}
				if rows != nil {
					DrawTableSepColBased(rows, cols)
				} else {
					ErrorHandler[60]()
					return
				}
			} else {
				ErrorHandler[20]()
				return
			}

		} else {
			ErrorHandler[80]()
		}
	}, // Get cast's saved networks
	"cast-devscan": func() {

	}, // Get cast to scan for wifi networks
	"cast-devforget": func() {
		if TargetMain != "" && WPAID != "" {
			url := PathsCast["forgetwifi"]
			newurl := fmt.Sprintf(url, TargetMain, DevicePorts["https"])
			JD := map[string]string{
				"wpa_id": WPAID,
			}
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			MakePost(newurl, tr, JD)
		} else {
			ErrorHandler[80]()
		}

	}, // Get cast to forget a network
	"cast-devrename": func() {
		if TargetMain != "" && DevName != "" {
			uri := PathsCast["setname"]
			newurl := fmt.Sprintf(uri, TargetMain, DevicePorts["https"])
			jsonData := map[string]string{
				"Content-Type": "application/json",
				"name":         DevName,
			}
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			MakePost(newurl, tr, jsonData)
		} else {
			ErrorHandler[80]()
		}
	}, // Get cast to rename itself
	"cast-devreboot": func() {
		if TargetMain != "" {
			uri := PathsCast["reset"]
			newurl := fmt.Sprintf(uri, TargetMain, DevicePorts["https"])
			JD := map[string]string{
				"params": "now",
			}
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			MakePost(newurl, tr, JD)
		} else {
			ErrorHandler[80]()
		}
	}, // Get cast to reboot itself
	"cast-devfreset": func() {
		if TargetMain != "" {
			uri := PathsCast["reset"]
			newurl := fmt.Sprintf(uri, TargetMain, DevicePorts["https"])
			JD := map[string]string{
				"params": "fdr",
			}
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			MakePost(newurl, tr, JD)
		} else {
			ErrorHandler[80]()
		}
	}, // Get cast to factory reset itself
	"cast-devkillapp": func() {}, // Get cast to kill any application
	"cast-devsetwall": func() {}, // Get cast to set wallpaper
	"cast-devplay":    func() {}, // Get cast to play videos
	"cast-devinfo": func() {
		if TargetMain != "" {
			url := fmt.Sprintf(Information1, TargetMain, DevicePorts["http"])
			resp, x := http.Get(url)
			if x != nil {
				ErrorHandler[10]()
				return
			}
			if resp.StatusCode == 200 {
				var ResponseSection0x1 SSDPXMLRESPONSEENTITY
				bv, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(bv, &ResponseSection0x1)
				var rows [][]string
				rows = append(rows,
					[]string{"Base URL", ResponseSection0x1.URLBase},
					[]string{"Device Type", ResponseSection0x1.Device.DeviceType},
					[]string{"Device Friendly Name", ResponseSection0x1.Device.FriendlyName},
					[]string{"Manufac", ResponseSection0x1.Device.Manufacturer},
					[]string{"Model Name", ResponseSection0x1.Device.ModelName},
					[]string{"UDN", ResponseSection0x1.Device.UDN},
					[]string{"Service Type", ResponseSection0x1.Device.ServiceList.Service.ServiceType},
					[]string{"Service ID", ResponseSection0x1.Device.ServiceList.Service.ServiceId},
					[]string{"Control URL", ResponseSection0x1.Device.ServiceList.Service.ControlURL},
					[]string{"EventSub URL", ResponseSection0x1.Device.ServiceList.Service.EventSubURL},
					[]string{"SCPDURL", ResponseSection0x1.Device.ServiceList.Service.SCPDURL},
				)
				DrawVerticle(rows)
			} else {
				ErrorHandler[20]()
			}
		} else {
			ErrorHandler[80]()
		}
	}, // Get cast to display deeper information about the device itself and hw
	"cast-info": func() {
		if TargetMain != "" {
			url := PathsCast["devinfo"]
			newurl := fmt.Sprintf(url, TargetMain, DevicePorts["http"])
			resp, x := http.Get(newurl)
			if x != nil {
				ErrorHandler[10]()
				return
			}
			if resp.StatusCode == 200 {
				var ResponseSection0x1 InfoResponseUnDetailed
				bv, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(bv, &ResponseSection0x1)
				var rows [][]string
				rows = append(rows,
					[]string{"BSSID", ResponseSection0x1.Bssid},
					[]string{"Build", ResponseSection0x1.BuildVersion},
					[]string{"Cast Build", ResponseSection0x1.CastBuildRevision},
					[]string{"Connected", fmt.Sprint(ResponseSection0x1.Connected)},
					[]string{"Ethernet Conn", fmt.Sprint(ResponseSection0x1.EthernetConnected)},
					[]string{"HasUpdate", fmt.Sprint(ResponseSection0x1.HasUpdate)},
					[]string{"HotSpotBSSID", fmt.Sprint(ResponseSection0x1.HotspotBssid)},
					[]string{"IPAddress", fmt.Sprint(ResponseSection0x1.IPAddress)},
					[]string{"Locale", fmt.Sprint(ResponseSection0x1.Locale)},
					[]string{"Location LAT", fmt.Sprint(ResponseSection0x1.Location.Latitude)},
					[]string{"Location LON", fmt.Sprint(ResponseSection0x1.Location.Longitude)},
					[]string{"Location CN", fmt.Sprint(ResponseSection0x1.Location.CountryCode)},
					[]string{"MAC", fmt.Sprint(ResponseSection0x1.MacAddress)},
					[]string{"name", fmt.Sprint(ResponseSection0x1.Name)},
					[]string{"SETUP History Succeed", fmt.Sprint(ResponseSection0x1.SetupStats.HistoricallySucceeded)},
					[]string{"SETUP Check connection", fmt.Sprint(ResponseSection0x1.SetupStats.NumCheckConnectivity)},
					[]string{"SETUP Connected Wifi networks", fmt.Sprint(ResponseSection0x1.SetupStats.NumConnectWifi)},
					[]string{"SETUP Connected Unsaved", fmt.Sprint(ResponseSection0x1.SetupStats.NumConnectedWifiNotSaved)},
					[]string{"SETUP Obtain IP", fmt.Sprint(ResponseSection0x1.SetupStats.NumObtainIP)},
					[]string{"SSDPUDN", fmt.Sprint(ResponseSection0x1.SsdpUdn)},
					[]string{"SSID", fmt.Sprint(ResponseSection0x1.Ssid)},
					[]string{"Time format", fmt.Sprint(ResponseSection0x1.TimeFormat)},
					[]string{"Uptime", fmt.Sprint(ResponseSection0x1.Uptime)},
					[]string{"Version", fmt.Sprint(ResponseSection0x1.Version)},
					[]string{"WPA Configured", fmt.Sprint(ResponseSection0x1.WpaConfigured)},
					[]string{"WPA State", fmt.Sprint(ResponseSection0x1.WpaState)},
				)
				DrawVerticle(rows)
			} else {
				ErrorHandler[20]()
			}

		} else {
			ErrorHandler[80]()
		}

	}, // Get cast information
	"roku-activeapp": func() {
		if TargetMain != "" {
			response, x := http.Get(fmt.Sprintf(PathsRoku["devactive"], TargetMain))
			if x != nil {
				fmt.Println("Error: Could not make a request to the device -> ", x)
			}
			if response.StatusCode == 200 {
				var Applications ActiveApp
				bv, _ := ioutil.ReadAll(response.Body)
				xml.Unmarshal(bv, &Applications)
				cols := []string{"AppID", "ApplicationName"}
				var (
					rows [][]string
				)
				for i := 0; i < len(Applications.App); i++ {
					rows = append(rows, []string{
						Applications.App[i].ID,
						Applications.App[i].Text,
					})
				}
				DrawTableSepColBased(rows, cols)
			} else {
				ErrorHandler[20]()
			}
		} else {
			ErrorHandler[80]()
		}
	}, // Get Roku current active app
	"roku-appinfo": func() {
		if TargetMain != "" {
			resp, x := http.Get(fmt.Sprintf(PathsRoku["devapps"], TargetMain))
			if x != nil {
				fmt.Println("Error when making request for APP information -> ", x)
				return
			}
			if resp.StatusCode == 200 {
				var Applications Apps
				bv, _ := ioutil.ReadAll(resp.Body)
				xml.Unmarshal(bv, &Applications)
				cols := []string{"Application ID", "Application Name"}
				rows := [][]string{}
				for i := 0; i < len(Applications.App); i++ {
					rows = append(rows, []string{
						Applications.App[i].ID,
						Applications.App[i].Text,
					})
				}
				DrawTableSepColBased(rows, cols)
			}
		} else {
			ErrorHandler[80]()
		}
	}, // Get Roku all current application and their ID's
	"roku-devstart": func() {
		if ApplicationIDROKU != "" && TargetMain != "" {
			resp, x := http.Get(fmt.Sprintf(PathsRoku["devapps"], TargetMain))
			if x != nil {
				fmt.Println("Error when making request for APP information -> ", x)
				return
			}
			var Applications Apps
			var Appname []string
			if resp.StatusCode == 200 {
				bv, _ := ioutil.ReadAll(resp.Body)
				xml.Unmarshal(bv, &Applications)
				for i := 0; i < len(Applications.App); i++ {
					Appname = append(Appname, strings.ToLower(string(Applications.App[i].ID)))
				}
			}
			url := Keys["launch"]
			for k := 0; k < len(Appname); k++ {
				if strings.Compare(ApplicationIDROKU, Appname[k]) == 0 {
					newurl := fmt.Sprintf(url, TargetMain, ApplicationIDROKU)
					NewPostNoData(newurl, true)
				}
			}
		} else {
			ErrorHandler[80]()
			ErrorHandler[90]()
		}
	}, // Get Roku to start an application based on its ID
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
			} else {
				ErrorHandler[20]()
			}
		}

	}, // Get RoKu information
	"roku-search": func() {
		if SearchQueryRoku != "" && TargetMain != "" {
			url := PathsRoku["browse"]
			newurl := fmt.Sprintf(url, TargetMain, SearchQueryRoku)
			NewPostNoData(newurl, true)
		} else {
			ErrorHandler[80]()
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
			ErrorHandler[80]()
		}

	}, // Get Roku to go home
	"roku-tvinfo": func() {
		if TargetMain != "" {
			url := Keys["tv"]
			url = fmt.Sprintf(url, TargetMain)
			f, x := http.Get(url)
			if x != nil {
				ErrorHandler[20]()
				return
			}
			defer f.Body.Close()
			var Byter TvChannels
			bv, _ := ioutil.ReadAll(f.Body)
			_ = xml.Unmarshal(bv, &Byter)
			Cols := []string{"Number", "Name", "Type", "User Hidden"}
			var rows [][]string
			for i := 0; i < len(Byter.Channel); i++ {
				rows = append(rows,
					[]string{
						Byter.Channel[i].Number,
						Byter.Channel[i].Name,
						Byter.Channel[i].Type,
						Byter.Channel[i].UserHidden,
					},
				)
			}
			if rows != nil {
				DrawTableSepColBased(rows, Cols)
			} else {
				ErrorHandler[60]()
				return
			}
		} else {
			ErrorHandler[80]()
		}
	}, // Get Roku TV information
	"roku-disable_sgr": func() {
		if TargetMain != "" {
			url := Keys["disable_sgr"]
			url = fmt.Sprintf(url, TargetMain)
			NewPostNoData(url, true)
			fmt.Println("[Information] SGR should be disabled")
		} else {
			ErrorHandler[80]()
		}
	}, // Disable Roku sgrendezvous
	"roku-enable_sgr": func() {
		if TargetMain != "" {
			url := Keys["enable_sgr"]
			url = fmt.Sprintf(url, TargetMain)
			NewPostNoData(url, true)
			fmt.Println("[Information] SGR should be enabled")
		}
	}, // Enable Roku sgrendezvous
	"roku-activetv": func() {
		if TargetMain != "" {
			url := Keys["activetv"]
			url = fmt.Sprintf(url, TargetMain)
			resp, x := http.Get(url)
			if x != nil {
				ErrorHandler[20]()
				return
			}
			defer resp.Body.Close()
			bv, _ := ioutil.ReadAll(resp.Body)
			var values TvChannel
			x = xml.Unmarshal(bv, &values)
			if x != nil {
				ErrorHandler[70]()
				fmt.Print(x)
				return
			}
			var rows [][]string
			rows = append(rows,
				[]string{"Chan Number", values.Channel.Number},
				[]string{"Chan Name", values.Channel.Name},
				[]string{"Chan Type", values.Channel.Type},
				[]string{"Chan UserHidden", values.Channel.UserHidden},
				[]string{"Active Input", values.Channel.ActiveInput},
				[]string{"Signal State", values.Channel.SignalState},
				[]string{"Signal Mode", values.Channel.SignalMode},
				[]string{"Signal quality", values.Channel.SignalQuality},
				[]string{"Signal strength", values.Channel.SignalStrength},
				[]string{"Program title", values.Channel.ProgramTitle},
				[]string{"Program Desc", values.Channel.ProgramDescription},
				[]string{"Program Rate", values.Channel.ProgramRatings},
				[]string{"Program Analog", values.Channel.ProgramAnalogAudio},
				[]string{"Program Digital", values.Channel.ProgramDigitalAudio},
				[]string{"Program Audio Formats", values.Channel.ProgramAudioFormats},
				[]string{"Program Audio Languages", values.Channel.ProgramAudioLanguages},
				[]string{"Program Audio Language", values.Channel.ProgramAudioLanguage},
			)
			if rows != nil {
				DrawVerticle(rows)
			} else {
				ErrorHandler[60]()
				return
			}
		} else {
			ErrorHandler[80]()
		}
	},
	"roku-getwalls": func() {
		if TargetMain != "" {
			resp, x := http.Get(fmt.Sprintf(PathsRoku["devapps"], TargetMain))
			if x != nil {
				fmt.Println("Error when making request for APP information -> ", x)
				return
			}
			if resp.StatusCode == 200 {
				var Applications Apps
				bv, _ := ioutil.ReadAll(resp.Body)
				xml.Unmarshal(bv, &Applications)
				cols := []string{"Application ID", "Application Name", "command to set appid"}
				rows := [][]string{}
				for i := 0; i < len(Applications.App); i++ {
					for _, k := range strings.Split(Applications.App[i].Text, " ") {
						if strings.Compare(strings.TrimSpace(k), "Screensaver") == 0 || strings.Compare(strings.TrimSpace(k), "Scrensaver") == 0 {
							rows = append(rows, []string{
								Applications.App[i].ID,
								Applications.App[i].Text,
								fmt.Sprintf("set appid=%s", Applications.App[i].ID),
							})
						}
					}
				}
				DrawTableSepColBased(rows, cols)
			}
		} else {
			ErrorHandler[80]()
		}
	},
	"roku-changewalls": func() {
		if TargetMain != "" && ApplicationIDROKU != "" {
			resp, x := http.Get(fmt.Sprintf(PathsRoku["devapps"], TargetMain))
			if x != nil {
				fmt.Println("Error when making request for APP information -> ", x)
				return
			}
			if resp.StatusCode == 200 {
				var Applications Apps
				bv, _ := ioutil.ReadAll(resp.Body)
				xml.Unmarshal(bv, &Applications)
				cols := []string{"Application ID", "Application Name"}
				rows := [][]string{}
				var iscorrect bool
				for i := 0; i < len(Applications.App); i++ {
					for _, k := range strings.Split(Applications.App[i].Text, " ") {
						if strings.Compare(strings.TrimSpace(k), "Screensaver") == 0 || strings.Compare(strings.TrimSpace(k), "Scrensaver") == 0 {
							if Applications.App[i].ID == ApplicationIDROKU {
								iscorrect = true
								rows = append(rows, []string{
									Applications.App[i].ID,
									Applications.App[i].Text,
								})
							}
							if iscorrect {
								fmt.Println(">>>  | Changing screensaver -> ", ApplicationIDROKU)
							}
						}
					}
				}
				if !iscorrect {
					fmt.Println("Error: Sorry could not find the screensaver with the ID you gave -> ", ApplicationIDROKU)
					return
				}
				if rows != nil {
					DrawTableSepColBased(rows, cols)
				} else {
					fmt.Println("Error: Sorry could not find the screensaver with the ID you gave -> ", ApplicationIDROKU)
					return
				}
				url := Keys["launch"]
				newurl := fmt.Sprintf(url, TargetMain, ApplicationIDROKU)
				NewPostNoData(newurl, true)
			}
		} else {
			ErrorHandler[80]()
		}
	},
}
