package CastHunter

import (
	"fmt"
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
	"roku-devscan":   func() {},
	"roku-devactive": func() {},
	"roku-devbrowse": func() {},
	"roku-up":        func() {},
	"roku-reboot": func() {
		for _, h := range RokuHosts {
			fmt.Print("\n")
			go func(h string) {
				if h == "10.0.0.59" {
					return
				}
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
	"roku-factory": func() {},
}
