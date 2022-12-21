package main

import (
	Caster "Casted/Modules"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

var (
	CFT = pflag.FlagSet{
		SortFlags: false,
	}
	CTFS Caster.Flags
)

func init() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	Caster.Banner()
	CFT.BoolVarP(&CTFS.Arp, "arp", "a", false, "Tell Caster's to send out an arp request to every device on the network silently in the background, then show hosts when you need to view them | This will locate chromecast devices on the network")
	CFT.IntVarP(&CTFS.SendOutEvery, "send", "e", 20, "Tell Caster's ARP module to send new ARP packets to every device on the network every so and so seconds, defualt is 20 seconds as it runs in the background")
	CFT.BoolVarP(&CTFS.OneInterface, "single", "s", true, "Tells Caster's ARP module to only work on one interface rather than all useable interfaces (default=true)")
	CFT.Parse(os.Args[1:])
	if CTFS.Arp {
		Caster.ArpActive = true
		go func() {
			Caster.Hunter(CTFS.OneInterface) // start ARP hunter
		}()
	}
	if CTFS.SendOutEvery != 0 {
		Caster.SleepInterval = CTFS.SendOutEvery
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	var Payload string
	time.Sleep(1 * time.Second)
	fmt.Print("\n\n\033[38;5;15mCaster> ")
	for {
		Payload, _ = input.ReadString('\n')              // read input until new line
		Payload = strings.Replace(Payload, "\n", "", -1) // read and replace input state
		if len(Payload) != 0 || len(Payload) != -1 {
			if Payload != "" {
				strindex := strings.Index(Payload, " ")
				if strindex == -1 {
					continue
				}
				switch Payload[0:strindex] {
				case "view":
					splitter := strings.Split(Payload[strindex:], " ")
					if splitter[1] != "" {
						if Caster.View[strings.TrimSpace(splitter[1])] != nil {
							Caster.View[strings.TrimSpace(splitter[1])]()
						} else {
							fmt.Println("[-] Not a value in the modules")
						}
					} else {
						fmt.Println("Value empty")
					}
				case "check":
					splitter := strings.Split(Payload[strindex:], " ")
					if splitter[1] != "" {
						if Caster.Check[strings.TrimSpace(splitter[1])] != nil {
							Caster.Check[strings.TrimSpace(splitter[1])]()
						} else {
							fmt.Println("[-] Not a value in the module")
						}
					} else {
						fmt.Println("[-] Sorry looks like you did not input a payload")
					}
				case "enumerate":
				case "set":
				default:
					fmt.Println("[-] Module did not exist, modules are -> (check, view, enumerate, set, help)")
				}
			} else {
				var help = "\t\tWelcome to Caster! the google cast hunter and enumerator! \n"
				help += "This framework allows you to hunt down and manipulate google Chrome Cast devices on the network \n"
				help += "below you will find a list of all commands, guides etc to having fun when exploiting your device \n"
				help += "------------------------------------------------------------------------------------------------"
				fmt.Println(help)
				fmt.Println("Syntax of module [set]       -> set variable=value | set target=1.1.1.1")
				fmt.Println("Syntax of module [enumerate] -> enumerate endpoint | enumerate devinfo")
				fmt.Println("Syntax of module [check]     -> check module       | check arp")
				fmt.Println("Syntax of module [view]      -> view variable      | view hosts")
				cols := []string{"Module", "Command", "Description"}
				rows := [][]string{}
				rows = append(rows,
					[]string{"set", "target", "Sets the target IP address of the chromecast device"},
					[]string{"check", "arp", "Checks if the ARP module is running"},
					[]string{"view", "hosts", "In a nice table format outputs the discovered hosts"},
					[]string{"view", "casts", "In a nice table format outputs the discovered chrome cast devices"},
					[]string{"enumerate", "devinfo", "Outputs device information of the set IP address / chromecast device"},
					[]string{"enumerate", "ipinfo", "Outputs the info of the set IP using an external API / chromecast devive"},
					[]string{"enumerate", "devsaved", "Outputs all wifi networks the chromecast device has saved"},
					[]string{"enumerate", "devscan", "Forces the set chromecast to scan for all wifi networks around"},
					[]string{"enumerate", "devforget", "Forces the set chromecast to forget a wifi network"},
					[]string{"enumerate", "devrename", "Forces the set chromecast to rename itself to a given name"},
					[]string{"enumerate", "devreboot", "Forces the device to reboot"},
					[]string{"enumerate", "devfreset", "Forces the device to factory reset"},
					[]string{"enumerate", "devkillapp", "Forces the device to kill a given application"},
				)
				Caster.DrawTableSepColBased(rows, cols)
			}
		}
		fmt.Print("Caster> ")
	}
}
