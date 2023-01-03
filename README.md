<p align="center">
  <img src="Caster.png">
</p>

# What is caster 

Caster is a IoT enumeration framework designed to abuse system API's and other forms of endpoints by constantly enumerating and searching for devices on the network. Caster can do many things depending on the device you are trying to enumerate and has many features which make it worth the use like the following below.

* Server: Once you start caster a local HTTP server will be started which outputs amazing information into different categories such as routers, devices and even gives you a good documentation in case you need help. ( Server is threaded seperately and can run offline )
* Port scanner: Because of issues that may occure with the ARP module figuring out how to categorize devices, caster has a built in port scanner which can categorize devices based on open static ports and ID's. This is something you can also use!
* ARP module: This module is the best part of caster as it allows you to discover devices on the network that can be categorized, either way by default this module runs in the background on a thread, it is recommended to keep this running if you want caster to run with full support
* Enumeration options: Caster comes with so much options and modules which one of the modules allows you to target every device caster has found at once. For example if caster has found 11 roku TV's on the same network and has verified their endpoint, you can use `enumerate* roku-reboot` which will start a thread for every host ( in this case 11 ) and attack all of the roku boxes at the same time. All roku boxes will most likely be rebooted.  

# Support 

Caster has support for many devices, some devices may be able to be controlled more than others such as roku's however there is still info and enumerating support for other devices. A list below provides this detail

| Device Name | Manufac | Function and Proc | 
| ----------- | ------- | ----------------- |
| FireTV      | Amazon  |  Device info      |
| Google Cast | Google  | Sys functions, System information, ID information, Kill functions etc |
| Cast Dongel | Google  | Sys functions, System information, ID information, Kill functions etc |
| Cast 4K     | Google  | Sys functions, System information, ID information, Kill functions etc |
| Roku mini   | Roku    | Sys functions, System information, ID information, Kill functions etc |
| Roku box    | Roku    | Sys functions, System information, ID information, Kill functions etc |
| AT&T STD Router | AT&T |  Device information |
| Apple TV    | Apple   |  Downloading informational files in apples binary format |
| CMS Samsung TV's | Samsung | Exploiting LFI, Exploiting system calls, device info |
| Rasberry PI's | Raspberry Pi Foundation | Scanning for SSh and other port forms |

# Install and Run 

* `sudo apt-get install libpcap-dev`
* `sudo apt-get install golang`
* `sudo go build -o caster main.go`
* `sudo ./caster [FLAGS]`

The word FLAGS in brackets you would replace with the following flags 

|Flag name and lable| Flag detail and info |
| ----------------- | -------------------- |
| --arp / -a        | Tells caster to run its ARP module ( suggested ) |
| --single / -s     | Tells caster's ARP module to only use a single interface, use this flag with --single=true to set this flag to true and use only a single interface, if this is set to false it can cause problems such as disconecting you from your network because caster will take every interface on the machine other than LO and ETH then proceed to use it for discovery |
| --send / -e | Tells caster's ARP module to send out a batch of packets every so and so seconds, the defualt is 20. For example if you set this flag using --send=30 it will send 255 ARP packets ( one for each host on the network ) every 30 seconds. so if you let caster run for 60 seconds with --send=30 caster would have sent 510 ARP packets |
| --server / -l | Tells caster to run a server in the background, this server is accessible via http://localhost:5429 and will display information such as documentations, error systems, devices organized, enumeration steps and docs on what the tool does, set this to false with --server=false if you do not want to use the server.The server will run on the background as a background thread / threaded process |
| --trace / -t | Tells casters error system and catching system to log and debug panics that happen within the program but will prevent the framework from crashing. These panics are not 100% guarenteed and sometimes it can be annoying, this is not due to a bug in the program but the way cards may be used or data may be sent or recieved |
| --errors / -r | Tells casters error system to display the entire error and throw a panic whenever there is a panic within the program |

* note: Do note that when you use --server the server indexes will refresh given the amount of time you tell caster to re send out packets, if you set the timer with --send=30 and use flag --server=true the server will regenerate every 30 seconds to refresh the data in  the indexes. The standard wait time for caster is 20 seconds which means the server will refresh every 20 seconds.

