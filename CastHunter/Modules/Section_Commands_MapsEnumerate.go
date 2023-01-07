package CastHunter

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Enum func()

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
	"roku-load": func() {
		RokuHosts = AddToList(RokuHosts, "Roku, Inc")
		RokuHosts = ExterminateExtraVals(RokuHosts)
	}, // Load all hosts and test
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
}
