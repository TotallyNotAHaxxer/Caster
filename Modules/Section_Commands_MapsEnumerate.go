package CastHunter

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var RokuHosts, CastHosts, FireTVHosts, AppleHosts []string
var RemAddr string

func AddToList(list []string, macstr string) []string {
	for i := 0; i < len(IPS); i++ {
		defer func() {
			if x := recover(); x != nil {
				return
			}
		}()
		mac := GrabManufac(MACS[i])
		if mac == macstr {
			list = append(list, IPS[i]) // Append the IP to the list
		}
	}
	if list != nil {
		return list
	} else {
		return nil
	}
}

func SleepLoop(t time.Duration) {
	time.Sleep(t * time.Millisecond)
}

var EnumerateAllMaps = map[string]Enum{
	"roku-poff": func() {
		for _, h := range RokuHosts {
			go func(h string) {
				url := Keys["devdown"]
				ur := fmt.Sprintf(url, h)
				NewPostNoData(ur, false)
			}(h)
		}
		time.Sleep(1 * time.Second)
	},
	"roku-pon": func() {
		for _, h := range RokuHosts {
			go func(h string) {
				NewPostNoData(fmt.Sprintf(Keys["devup"], h), false)
			}(h)
		}
		time.Sleep(1 * time.Second)
	},
	"roku-load": func() {
		RokuHosts = AddToList(RokuHosts, "Roku, Inc")
		RokuHosts = ExterminateExtraVals(RokuHosts)
	}, // Load all hosts and test
	"roku-console": func() {
		RokuDevStartConsoleall()
	},
	"roku-devstart": func() {
		if ApplicationIDROKU != "" {
			for _, h := range RokuHosts {
				go func(h string) {
					OK[99]()
					SApp(h)
				}(h)
			}
			SleepLoop(1000)
		}
	},
	"roku-devhome": func() {
		for _, h := range RokuHosts {
			go func(h string) {
				OK[99]()
				Home(h)
			}(h)
		}
	},
	"roku-devscan": func() {},
	"roku-devactive": func() {
		colums := []string{"Hostname", "Active App on Host", "APP ID"}
		var rows [][]string
		for _, l := range RokuHosts {
			go func(l string) {
				url := PathsRoku["devactive"]
				ur := fmt.Sprintf(url, l)
				f, x := http.Get(ur)
				if x != nil {
					ErrorHandler[10]()
					fmt.Print(x)
					return
				}
				if f.StatusCode == 200 {
					var res ActiveApp
					bv, _ := ioutil.ReadAll(f.Body)
					xml.Unmarshal(bv, &res)
					for i := 0; i < len(res.App); i++ {
						rows = append(rows, []string{l, res.App[i].Text, res.App[i].ID})
					}
				}
			}(l)
		}
		time.Sleep(5 * time.Second) // give it some time to finish
		if rows != nil {
			DrawTableSepColBased(rows, colums)
		} else {
			ErrorHandler[150]()
			return
		}
	},
	"roku-devbrowse": func() {
		if SearchQueryRoku != "" {
			for _, l := range RokuHosts {
				go func(l string) {
					url := PathsRoku["browse"]
					ur := fmt.Sprint(url, l, SearchQueryRoku)
					NewPostNoData(ur, false)
				}(l)
			}
		}
	},
	"roku-up": func() {
		for _, h := range RokuHosts {
			go func(h string) {
				Up(h)
			}(h)
		}
	},
	"roku-reboot": func() {
		for _, h := range RokuHosts {
			fmt.Print("\n")
			go func(h string) {
				OK[99]()
				ApplicationIDROKU = "2285"
				SApp(h)
				ApplicationIDROKU = ""
				OK[98]()
				SleepLoop(5000)
				Home(h)
				SleepLoop(3000)
				Up(h)
				SleepLoop(400)
				Click(h)
				for i := 0; i < 12; i++ {
					Down(h)
					SleepLoop(400)
				}
				SleepLoop(400)
				Right(h)
				SleepLoop(400)
				for k := 0; k < 7; k++ {
					Down(h)
					SleepLoop(400)
				}
				Click(h)
				SleepLoop(400)
				Click(h)
			}(h)
		}
	},
	"roku-factory": func() {}, // as of now can not do this
	// google cast enumerate all
	"cast-rename": func() {
		if DevName == "" {
			input := bufio.NewReader(os.Stdin)
			fmt.Println("\x1b[33m (WARN): In order to rename every device that is a google cast on the network Caster needs a VALID AND NOT NULL name to be set, prompting user input...")
			fmt.Print("Name to reset cast to> ")
			for {
				var Payload string
				Payload, _ = input.ReadString('\n')
				Payload = strings.Replace(Payload, "\n", "", -1)
				if Payload != "" {
					DeviceRaname = Payload
					break
				} else {
					fmt.Print("Sorry this was not a valid payload, it was empty, try again")
				}
				fmt.Print("Name to reset cast to> ")
			}
		}
		fmt.Println("Renaming device to -> ", DevName)
		jsonData := map[string]string{
			"Content-Type": "application/json",
			"name":         DeviceRaname,
		}
		for _, h := range GoogleHost_Enum {
			go func(h string) {
				uri := PathsCast["setname"]
				newurl := fmt.Sprintf(uri, h, DevicePorts["https"])
				MakePost(newurl, tr, jsonData, true)
			}(h)
		}
		fmt.Println("[DEBUG](INFO): Sleeping threads before returning to input, TIME WAIT...")
		time.Sleep(3 * time.Second)
	},

	"cast-reboot": func() {
		for _, h := range GoogleHost_Enum {
			go func(h string) {
				uri := PathsCast["reset"]
				newurl := fmt.Sprintf(uri, h, DevicePorts["https"])
				MakePost(newurl, tr, JD_Params, true)
			}(h)
		}

	},
	"cast-factoryreset": func() {
		for _, h := range GoogleHost_Enum {
			go func(h string) {
				uri := PathsCast["reset"]
				newurl := fmt.Sprintf(uri, h, DevicePorts["https"])
				MakePost(newurl, tr, JD_Paramsfr, true)
			}(h)
		}
	},
	// Entry is an auto run function, if entry is not used no hosts will be loaded, it is also proper to ensure all duplicates are removed
	"entry": func() {
		for i := 0; i < len(MACS); i++ {
			if len(MACS) == len(IPS) {
				// continue to write the table
				manufac := GrabManufac(MACS[i])
				if manufac == "Google, Inc." {
					GoogleHost_Enum = append(GoogleHost_Enum, IPS[i])
				} else {
					cl := http.Client{
						Timeout: time.Second * 2,
					}
					host := IPS[i]
					go func(host string) {
						response, x := cl.Get(fmt.Sprintf(PathsCast["devinfo"], host, "8008"))
						if x == nil {
							if response.StatusCode == 200 {
								GoogleHost_Enum = append(GoogleHost_Enum, host)
							}
						}
					}(host)
				}
			}
		}
		fmt.Println("[Information] Sleeping threads for (5) seconds, WAIT TIME FOR FINISH (CHANS)")
		time.Sleep(5 * time.Second)
	},
}
