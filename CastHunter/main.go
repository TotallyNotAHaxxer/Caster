package main

// TODO: Re-do this entire file on further versions, atomic operations are just copied from themselves and just changed with one number, this can cause performance issues in the future, just ensure we fix it ~ Totally_Not_A_Haxxer

import (
	Caster "Casted/Modules"
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
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

func CheckLog() bool {
	if CTFS.Trace {
		return true
	}
	return false
}

func CheckMessages(x interface{}) {
	if CTFS.Trace {
		fmt.Println("\033[31m", x)
	}
}

func init() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	Caster.Banner()
	CFT.BoolVarP(&CTFS.Arp, "arp", "a", false, "Tell Caster's to send out an arp request to every device on the network silently in the background, then show hosts when you need to view them | This will locate chromecast devices on the network")
	CFT.IntVarP(&CTFS.SendOutEvery, "send", "e", 20, "Tell Caster's ARP module to send new ARP packets to every device on the network every so and so seconds, defualt is 20 seconds as it runs in the background")
	CFT.BoolVarP(&CTFS.OneInterface, "single", "s", true, "Tells Caster's ARP module to only work on one interface rather than all useable interfaces (default=true)")
	CFT.BoolVarP(&CTFS.Server, "server", "l", true, "Tell caster to load a HTTP server in the background for documentation")
	CFT.BoolVarP(&CTFS.Trace, "trace", "t", false, "Tells Caster's error system to spit out panic responses messages when the program attempts to crash")
	CFT.BoolVarP(&CTFS.Error, "errors", "r", false, "Tells Caster to throw panic errors when they are caught")
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
		// the trashiest way of doing things-
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
		timer := time.NewTicker(time.Second * time.Duration(CTFS.SendOutEvery))
		var Atom atomic.Value
		Atom.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",

				fmt.Sprint(len(Caster.IPS)), "22",
				fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery),
			)
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					if mac == "Roku, Inc" {
						Doc += "<th>" + Caster.IPS[i] + "</th>"
						Doc += "<th>" + Caster.MACS[i] + "</th>"
						Doc += "<th>" + mac + "</th>"
					}
					Doc += "</tr>"
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/roku.html"
			os.Remove(file)
			_ = ioutil.WriteFile(file, []byte(Doc), 777)
		})
		go func() {
			for {
				select {
				case <-timer.C:
					fun := Atom.Load().(func())
					fun()
				}
			}
		}()
		var Atom2 atomic.Value
		Atom2.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				fmt.Sprint(len(Caster.IPS)), "22", fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery))
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					if mac == "Unknown" || mac == "Could not be determined | Length issue in array (PANIC)" {
						cl := http.Client{
							Timeout: time.Duration(1 * time.Second),
						}
						resp, x := cl.Get(fmt.Sprintf("http://%s:8008/setup/eureka_info", Caster.IPS[i]))
						if x == nil {
							Doc += "<th>" + Caster.IPS[i] + "</th>"
							Doc += "<th>" + Caster.MACS[i] + "</th>"
							Doc += "<th>" + mac + " Found with PORT scan</th>"
							defer resp.Body.Close()
						}
					}

					if mac == "Google, Inc." {
						Doc += "<th>" + Caster.IPS[i] + "</th>"
						Doc += "<th>" + Caster.MACS[i] + "</th>"
						Doc += "<th>" + mac + "</th>"
					}

					Doc += "</tr>"
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/cast.html"
			_ = ioutil.WriteFile(file, []byte(Doc), 777)

		})
		go func() {
			for {
				select {
				case <-timer.C:
					os.Remove("Modules/Static/cast.html")
					fun := Atom2.Load().(func())
					fun()
				}
			}
		}()
		var Atom3 atomic.Value
		Atom3.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				fmt.Sprint(len(Caster.IPS)), "22", fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery))
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					Doc += "<th>" + Caster.IPS[i] + "</th>"
					Doc += "<th>" + Caster.MACS[i] + "</th>"
					Doc += "<th>" + mac + "</th>"
					Doc += "</tr>"
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/devs.html"
			_ = ioutil.WriteFile(file, []byte(Doc), 777)
		})
		go func() {
			for {
				select {
				case <-timer.C:
					os.Remove("Modules/Static/devs.html")
					fun := Atom3.Load().(func())
					fun()

				}
			}
		}()
		var Atom4 atomic.Value
		Atom4.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				fmt.Sprint(len(Caster.IPS)), "0", fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery))
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					if mac == "Raspberry Pi Trading Ltd" {
						Doc += "<th>" + Caster.IPS[i] + "</th>"
						Doc += "<th>" + Caster.MACS[i] + "</th>"
						Doc += "<th>" + mac + "</th>"
						Doc += "</tr>"
					}
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/rpis.html"
			_ = ioutil.WriteFile(file, []byte(Doc), 777)
		})
		go func() {
			for {
				select {
				case <-timer.C:
					os.Remove("Modules/Static/rpis.html")
					fun := Atom4.Load().(func())
					fun()

				}
			}
		}()
		var Atom5 atomic.Value
		Atom5.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				fmt.Sprint(len(Caster.IPS)), "1", fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery))
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					if mac == "Apple, Inc." {
						Doc += "<th>" + Caster.IPS[i] + "</th>"
						Doc += "<th>" + Caster.MACS[i] + "</th>"
						Doc += "<th>" + mac + "</th>"
						Doc += "</tr>"
					}
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/adevs.html"
			_ = ioutil.WriteFile(file, []byte(Doc), 777)
		})
		go func() {
			for {
				select {
				case <-timer.C:
					os.Remove("Modules/Static/adevs.html")
					fun := Atom5.Load().(func())
					fun()

				}
			}
		}()
		var Atom6 atomic.Value
		Atom6.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				fmt.Sprint(len(Caster.IPS)), "0", fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery))
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					if mac == "Amazon Technologies Inc." {
						Doc += "<th>" + Caster.IPS[i] + "</th>"
						Doc += "<th>" + Caster.MACS[i] + "</th>"
						Doc += "<th>" + mac + "</th>"
						Doc += "</tr>"
					}
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/azdevs.html"
			_ = ioutil.WriteFile(file, []byte(Doc), 777)
		})
		go func() {
			for {
				select {
				case <-timer.C:
					os.Remove("Modules/Static/azdevs.html")
					fun := Atom6.Load().(func())
					fun()
				}
			}
		}()
		var Atom7 atomic.Value
		Atom7.Store(func() {
			var Doc string
			Doc += Caster.STDHTML
			Doc = fmt.Sprintf(Doc,
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				"active",
				"inactive",
				"inactive",
				"inactive",
				"inactive",
				fmt.Sprint(len(Caster.IPS)), "2", fmt.Sprint(len(Caster.MACS)), fmt.Sprint(CTFS.SendOutEvery))
			if Caster.IPS != nil {
				for i := 0; i < len(Caster.IPS); i++ {
					Doc += "<tr>"
					var mac string
					if Caster.CheckLen() {
						mac = Caster.GrabManufac(Caster.MACS[i])
					} else {
						defer func() {
							if x := recover(); x != nil {
								if CheckLog() {
									Caster.ErrorHandler[120]()
								}
								CheckMessages(x)
								return
							}
						}()
						mac = Caster.GrabManufac(Caster.MACS[i])
					}
					if mac == "ARRIS Group, Inc." {
						Doc += "<th>" + Caster.IPS[i] + "</th>"
						Doc += "<th>" + Caster.MACS[i] + "</th>"
						Doc += "<th>" + mac + "</th>"
						Doc += "</tr>"
					}
				}
			} else {
				Doc += "</tr>"
				Doc += "<th>IP Addresses to scan were NULL</th>"
				Doc += "<th>MAC Addresses to scan were NULL</th>"
				Doc += "<th>OUI's mined were NULL</th>"
				Doc += "</tr>"
			}
			Doc += Caster.STDHTMLBOTTOM
			f, x := os.Open("Modules/Templates/SameSideCSSBars.css.tmpl")
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Doc += scanner.Text()
				}
			}
			file := "Modules/Static/ro.html"
			_ = ioutil.WriteFile(file, []byte(Doc), 777)
		})
		go func() {
			for {
				select {
				case <-timer.C:
					os.Remove("Modules/Static/ro.html")
					fun := Atom7.Load().(func())
					fun()
				}
			}
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
				case "enumerate*":
					splitter := strings.Split(Payload[strindex:], " ")
					if splitter[1] != "" {
						if strings.Contains(splitter[1], "roku") {
							Caster.EnumerateAllMaps["roku-load"]()
						}
						if Caster.EnumerateAllMaps[strings.TrimSpace(splitter[1])] != nil {
							Caster.EnumerateAllMaps[strings.TrimSpace(splitter[1])]()
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
