package main

import (
	Caster "Casted/Modules"
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	CFT.BoolVarP(&CTFS.Server, "server", "l", true, "Tell caster to load a HTTP server in the background for documentation")
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
	if CTFS.Server {
		watch := time.Now()
		server := &http.Server{
			Addr: ":5429", // Address of server
		}
		go func() {
			fmt.Println("[Information] Server has been started on port 5429 (http://localhost:5429)")
			http.HandleFunc("/", Caster.Process_Incoming_ReqStream)
			server.ListenAndServe()
		}()
		go func() {
			sigwait := make(chan os.Signal, 1)
			signal.Notify(sigwait, syscall.SIGINT, syscall.SIGKILL)
			<-sigwait
			contect, cl := context.WithTimeout(context.Background(), 1*time.Second)
			defer cl()
			fmt.Println("\x1b[33m \n[Information] \x1b[38;5;13mServer shutdown at -> localhost@8080 ")
			fmt.Println("\x1b[33m[Information] \x1b[38;5;13mShutting down sub threads...")
			if x := server.Shutdown(contect); x != nil {
				fmt.Println("\033[31m[ERROR] For some reason the HTTP server started on -> http://localhost" + server.Addr)
				fmt.Println("[ERROR] Was not properly shut off, make sure to shut this server down as this can cause conflicts")
				fmt.Println("[ERROR] Within this program or other programs that may be hosting the server on the same port")
			} else {
				fmt.Println("\x1b[33m[Information] \x1b[38;5;13m Server has been shutdown correctly")
				fmt.Println("\x1b[92m[User Info  ] If the server was not shut down within the program it is suggested to shut it down yourself")
				fmt.Println("\x1b[92m[User Info  ] It may also be good to note to report this as a developer error")
				fmt.Println("\x1b[0m")
				var messsages string
				messsages += "Server uptime    => " + fmt.Sprint(time.Since(watch)) + "\n"
				messsages += "Server Port      => " + server.Addr + "\n"
				messsages += "Server Host      => http://localhost" + server.Addr
				Caster.DrawUtilsBox(messsages)
				os.Exit(0)
			}
			select {}
		}()
	}
}

func DisplayH() {
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
		[]string{"view", "cls", "Clears the screen"},
		[]string{"view", "clear", "Clears the screen"},
		[]string{"view", "artc", "Clears the screen"},
		[]string{"enumerate", "devinfo", "Outputs device information of the set IP address / chromecast device"},
		[]string{"enumerate", "ipinfo", "Outputs the info of the set IP using an external API / chromecast devive"},
		[]string{"enumerate", "devsaved", "Outputs all wifi networks the chromecast device has saved"},
		[]string{"enumerate", "devscan", "Forces the set chromecast to scan for all wifi networks around"},
		[]string{"enumerate", "devforget", "Forces the set chromecast to forget a wifi network"},
		[]string{"enumerate", "devrename", "Forces the set chromecast to rename itself to a given name"},
		[]string{"enumerate", "devreboot", "Forces the device to reboot"},
		[]string{"enumerate", "devfreset", "Forces the device to factory reset"},
		[]string{"enumerate", "devkillapp", "Forces the device to kill a given application"},
		[]string{"enumerate", "*ports", "Will scan all google selected or found devices"},
		[]string{"enumerate", "*hosts", "Will take all cast set hosts and run functions on all hosts set for casts"},
		[]string{"enumerate", "*run", "Will run every single function with default arguments on a single set host"},
		[]string{"enumerate", "ports", "Will scan the set target for enumeration for a list of common ports, any that have ports 8008, 8443, 8009 will be deemed as a cast"},
	)
	Caster.DrawTableSepColBased(rows, cols)
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
				if Payload == "help" || Payload == "h" || Payload == "he" || Payload == "hh" {
					DisplayH()
				}
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
					splitter := strings.Split(Payload[strindex:], " ")
					if splitter[1] != "" {
						if Caster.Enumerate[strings.TrimSpace(splitter[1])] != nil {
							Caster.Enumerate[strings.TrimSpace(splitter[1])]()
						}
					}
				case "set":
					splitter := strings.Split(Payload[strindex:], "=")
					if splitter[0] != "" && splitter[1] != "" {
						if Caster.Set[strings.TrimSpace(splitter[0])] != nil {
							Caster.Set[strings.TrimSpace(splitter[0])](strings.TrimSpace(splitter[1]))
						}
					}
				default:
					fmt.Println("[-] Module did not exist, modules are -> (check, view, enumerate, set, help)")
				}
			} else {
				DisplayH()
			}
		} else {
			fmt.Println("[-] Error: The command you entered may not exist, it was not of a good length, it was 0 or -1 in indexing...")
		}
		fmt.Print("Caster> ")
	}
}
