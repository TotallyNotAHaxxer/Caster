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
	return CTFS.Trace
}

func CheckMessages(x interface{}) {
	if CTFS.Trace {
		fmt.Println("\033[31m", x)
	}
}

func init() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	Caster.Banner()
	CFT.BoolVarP(&CTFS.Arp, "arp", "a", true, "Tell Caster's to send out an arp request to every device on the network silently in the background, then show hosts when you need to view them | This will locate chromecast devices on the network")
	CFT.IntVarP(&CTFS.SendOutEvery, "send", "e", 20, "Tell Caster's ARP module to send new ARP packets to every device on the network every so and so seconds, defualt is 20 seconds as it runs in the background")
	CFT.BoolVarP(&CTFS.OneInterface, "single", "s", false, "Tells Caster's ARP module to only work on one interface rather than all useable interfaces (default=true)")
	CFT.BoolVarP(&CTFS.Server, "server", "l", true, "Tell caster to load a HTTP server in the background for documentation")
	CFT.BoolVarP(&CTFS.Trace, "trace", "t", false, "Tells Caster's error system to spit out panic responses messages when the program attempts to crash")
	CFT.BoolVarP(&CTFS.Error, "errors", "r", false, "Tells Caster to throw panic errors when they are caught")
	CFT.BoolVarP(&CTFS.SSDP, "ssdp", "d", true, "Tells Caster to choose any interface to run a SSDP listener, (defualt = true) | Note: This setting is required to enumerate certain devices as a UUID or serial number of the device is required, if it is not, please pre set the UUID if you know the UUID of the device")
	CFT.BoolVar(&CTFS.Help, "help", false, "load help")
	CFT.BoolVar(&CTFS.SuperHelp, "shelp", false, "load deeper help menu")
	CFT.Parse(os.Args[1:])
	if CTFS.SuperHelp {
		Caster.Help2()
		os.Exit(0)
	}
	if CTFS.Help {
		Caster.HelpMenu()
		os.Exit(0)
	}
	if CTFS.Arp {
		Caster.ArpActive = true
		go func() {
			Caster.Hunter(CTFS.OneInterface) // start ARP hunter
		}()
	}
	if CTFS.SSDP {
		go func() {
			Caster.Listener(Caster.RetInterface().Name) // Start SSDP listener
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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)
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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)

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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)
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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)
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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)
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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)
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
			_ = ioutil.WriteFile(file, []byte(Doc), 0777)
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
					fmt.Println(`

						Caster V1.0 STABLE
------------------------------------------------------------
Welcome to caster, the ghost that hunts down IoT devices for 
Enumeration. This is the help menu for casters console. Below 
you will find a list of all of casters console based commands 
what they do and the requirements they may require as well as
a list of modules and their commands without descriptions.

Module                  Commands
----------             -----------
enumerate              ports|appleinfo|cast-devforget|cast-devfreset|
						cast-devinfo|cast-devreboot|cast-devrename|cast-devsaved|
						cast-info|firestick-devinfo|roku-activeapp|roku-activetv|
						roku-appinfo|roku-changewalls|roku-console|roku-devinfo|
						roku-devstart|roku-disable_sgr|roku-enable_sgr|roku-getwalls|
						roku-home|roku-search|roku-tvinfo|routerXF-GDD|samsungCMS-off

view                   amazon|apple|artc|casts|clear|cls|hosts|roku|routers|scannresults|uuids

enumerate*            cast-factoryreset|cast-reboot|cast-rename|roku-devactive|roku-devbrowse|
						roku-devhome|roku-devstart|roku-reboot|roku-up|roku-console

check                 arp|server| 

set                   target|keyword|appid|devname|wpaid

COMMAND DESCRIPTION, WALKTHROUGH AND UNDERSTANDING 

================================ MODULE (enumerate)

Enumerate is the biggest module in caster, it contains over 30 commands for 10 different 
devices ( some which are not fully implemented yet). It is important to note that while 
every single command was tested that when you use it, it may not work the same. What this 
means is that not all commands, endpoints, packets, payloads etc are accepted on every host.
for example if you try to use the enumerate cast-reboot it will not reboot on all google casts 
this was tested and found out that only certain versions of the casts paths are supported. However 
what is supported it information of the devices no matter the version, but running system functions 
is not always supported. So do beware of this as the program is not broken, vulnerabilities were 
not patched, and google or the device manufac's did not remove these endpoints. they just exist 
simply on other devices. 

MODULE SYNTAX 

enumerate (command)

REQUIRES: set target - in order to use any of the enumerate options you must SET a target 
for the current session ( changeable anytime ) using command 'set target' note that anytime 
you use the module 'set' to set a command variable you MUST put = after the variable. In our 
case if we wanted to set the target to 10.0.0.2 we would say 'set target=10.0.0.2'. This is 
where it gets easier. Once target is valid you can attack any device as long as its supported 
on the list. COMMANDS MUST BE BY DEVICE YOU CAN NOT RUN ROKU FUNCTIONS ON A GOOGLE CAST IT 
DOES NOT WORK LIKE THAT.

MODULE COMMANDS 

enumerate              ports|appleinfo|cast-devforget|cast-devfreset|
						cast-devinfo|cast-devreboot|cast-devrename|cast-devsaved|
						cast-info|firestick-devinfo|roku-activeapp|roku-activetv|
						roku-appinfo|roku-changewalls|roku-console|roku-devinfo|
						roku-devstart|roku-disable_sgr|roku-enable_sgr|roku-getwalls|
						roku-home|roku-search|roku-tvinfo|routerXF-GDD|samsungCMS-off

enumerate ports: runs a simple port scan on the host to verify open ports, only 
					will scan basic and important ports not all 65K ports 

enumerate appleinfo: runs a function which will extract a PLIST file from an APPLETV 
					host which will then disect and verify the file information.

enumerate cast-devforget: runs a function which will tell a google cast to forget a 
						network based on its WPAID

enumerate cast-devfreset: will attempt to factory reset a google cast 

enumerate cast-devinfo: will attempt to grab device information on a google cast 

enumerate cast-devreboot: will attempt to reboot a google cast 

enumerate cast-devrename: will attempt to rename the device given a set name 'set devname=devname'

enumerate cast-devsaved: will attempt to get saved wifi networks on the google cast 

enumerate cast-info: will attempt to grab device information on the google cast 

enumerate firestick-devinfo: will attempt to grab the device information of an amazonFireTV
NOTE: the amazon firstick uses a security feature which prevents you from accessing the information 
of a FireTV without having a valid UUID. Luckily amazonFireTV's also send out SSDP packets containing 
the information for its endpoints. Caster will sniff for UUID's and when running this function even if 
there is an amazon fireTV on the network if caster has not found a valid UUID for the host it will not 
run the function. If you need deeper knowledge on this please go to the article titled 
'Abusing API’s in IoT Devices' on medium at https://medium.com/@Totally_Not_A_Haxxer

enumerate roku-activeapp: Will attempt to grab the currently active application and its ID 
							on a rokuTV box 

enumerate roku-activetv: Will attempt to grab the most active TV channel on a rokuTV box 

enumerate roku-appinfo: Will attempt to grab ever application on a RokuBox along with its ID 

enumerate roku-changewalls: Will attempt to grab and change a wallpaper or screensaver on a RokuBox 

enumerate roku-console: This command will launch a fully interactive console for enumerating a RokuBox 
by sending remote controlls, this gives you the power of the RokuRemote on your computer as if you were 
holding it. This will be able to make the OS go left, right, up, down, home, launch apps, mute volume, 
turn volume up, turn volume down etc etc

enumerate roku-devinfo: will attempt to grab device information of a roku box 

enumerate roku-devstart: will attempt to start an application based on its ID selected or set within the 
						console
					
enumerate roku-disable_sgr: will disable SGR 

enumerate roku-enavble_sgr: will enable SGR

enumerate roku-getwalls: Will attempt to get all installed wallpapers or screensavers on the device installed 

enumerate roku-home: Will attempt to send the roku TV to home 

enumerate roku-search: Given a keyword it will attempt to make a RokuBOX go home and search for this topic 

enumerate roku-tvinfo: WIll get current TV information of a RokuBOX if TV is enabled

enumerate routerXF-GDD: Will attempt to grab router information or device information from a standard 
						Xfinity home router or any router that has ties with ARRIS GROUP

enumerate samsungCMS-off: Will attempt to exploit a specific SAMSUNG SMART TV by powering it off and taking 
							advantage of a vulnerability within the system service API

================================ MODULE (check)

check is a simple module, all its purpose is, is to check and verify the state of a service or thread 
within caster 

commands

check arp: will check if the arp module is running 
check server: will check if the local server is running 

================================ MODULE (set)

set is by far the second most important module if not up there to be first 
within caster. It is the reason you can enumerate certain devices. Set allows 
you to set script and session variables such as the target, appid, device name etc 

COMMANDS 

set                   target|keyword|appid|devname|wpaid

syntax: set variable=value

set target: Set target allows you to set a target to enumerate, this target is changeable anytime.
			Execute with set target=target.ip.address.4

set keyword: Set keyword allows you to set a search keyword for roku enumeration on search and browse 
				based functions. Execute this with --> set keyword=searchkeywordorquery

set appid: Set appid allows you to set the appid of the application you please to launch on roku boxes 
			Execute with - set appid=appidnumber

set devname: set devname allows you to set the device name you would like to change a google cast to 

set wpaid: set wpaid allows you to set the WPAID of a network you want a google cast device to forget 

=============================== MODULE (view)

this module is super light however it allows you to view 
data that caster has collected like hosts, UUID's, brand names etc.
A cool thing about caster along with its interfaces is in the event that 
you are on a network with a large amount of devices, it will seperate them 
per supported brand such as 

amazon, google, roku, arrisgroup, apple etc

the following commands exist for the view module 

view amazon: Will output all captured amazon devices
view casts: Will output all captured google cast devices 
view routers: Will attempt to output all found ARRISGROUP routers 
view roku: Will output all captured roku devices
view apple: will output all captured apple devices
view artc: will clear the screen and display the banner 
view cls: clears screen 
view clear: clears screen 
view hosts: Will output all found hosts by the ARP module 


view uuids: will output all found UUID's and their hosts found by the SSDP module



thats all to this module, pretty simple


================================ MODULE (enumerate*)

Enumerate* is the most powerfull module within this framework. This framework on the right network 
can raise hell or really cause a bunch of confusion. In a sense this module will grab every device 
that casters ARP module has found that is of a specific brand or supported device and execute a 
function that you choose. For example if caster has a table of Roku devices like so 


┣━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ IP Address ┃ MAC Address       ┃ Manufac                  ┃
┣━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┃ 10.0.0.0   ┃ ff:ff:ff:ff:ff:ff ┃ Roku, Inc                ┃
┣━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━━━━━━━━━━━┫

and you say 'enumerate* roku-home' caster will take every host in that list 
that has a manufac value of 'Roku, Inc' and will take those IP addresses and 
on a seperate thread for each host run that function on the host. So you would 
be sending over 12 devices to the home page.Enumerate all is a really interesting 
feature to caster and it even allows you to start an interactive console for roku 
devices and run the same controls in the other console 'enumerate roku-console' just for every single host

enumerate* is the same thing as enumerate with the same functions and no different, it 
should be pretty easy to understand.

COMMANDS 
	enumerate*            cast-factoryreset|cast-reboot|cast-rename|roku-devactive|roku-devbrowse|
							roku-devhome|roku-devstart|roku-reboot|roku-up|roku-console

					
					`)

					continue
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
						} else if strings.Contains(splitter[1], "cast") {
							Caster.EnumerateAllMaps["entry"]()
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
				fmt.Println("Type `help` if you do not know what you are doing or need a reminder of commands")
			}
		} else {
			fmt.Println("[-] Error: The command you entered may not exist, it was not of a good length, it was 0 or -1 in indexing...")
		}
		fmt.Print("Caster> ")
	}
}
