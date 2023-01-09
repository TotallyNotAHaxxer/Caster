package CastHunter

import (
	"fmt"
	"os/user"
)

// Error handler
type ErrorHandle func()

var ErrorsPackets = map[string]ErrorHandle{
	"pcapOpenLive": func() {
		fmt.Println("[--!!--] Error: There was an issue when trying to open up a new PCAP handler using the interface given, this may be due to permission issues. Is user sudo?")
		use, x := user.Current()
		if x != nil {
			fmt.Println("[--!FATAL!--] FATAL ERROR SYS: Could not manage to get the user, this is an issue when trying to check for the root user or user group -> ", x)
			return
		}
		ru, x := user.Lookup("root")
		if x != nil {
			fmt.Println("[--!FATAL!--] FATAL ERROR SYS: Could not manage to get the user, this is an issue when trying to check for the root user or user group -> ", x)
			return
		}
		if use.Uid == ru.Uid {
			fmt.Println("[Information] Info (USER PRIV DEBUG) => The user is root ( ", use.Name, " ) However there was still an error getting the live interface to open")
		} else {
			fmt.Println("[--!!--] Error: ( ", use.Name, " ) is NOT running the program as root, please ensure caster is run as user like so `sudo go run main.go` or running the binary as root so caster cann access drives and other external and internal resources")

		}
	},
	"filtere": func() {
		fmt.Println("[--!!--] Error: I could not seem to actually set the filter, this filter is required to capture packets but caster recieved an error, broken handler?")
	},
	"uuidnil": func() {
		fmt.Println("[--!!--] Error: Sorry it seems as if there are no current UUID's for devices that are on your target list or within the range. SSDP scanner on?")
	},
}
