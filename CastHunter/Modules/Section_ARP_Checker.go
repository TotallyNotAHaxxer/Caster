package CastHunter

import (
	"fmt"
	"net"
)

func CheckAddress(castaddr *net.IPNet) bool {
	if castaddr == nil {
		fmt.Printf("\n\033[38;5;202m[Warning] Skipping \033[38;5;50m[\033[38;5;15m%s\033[38;5;50m] \033[38;5;196m(False Network)", castaddr)
		return false
	} else if castaddr.IP[0] == 127 {
		fmt.Printf("\n\033[38;5;202m[Warning] Skipping \033[38;5;50m[\033[38;5;15m%s\033[38;5;50m] \033[38;5;196m(LocalHost)", castaddr)
		return false
	} else if castaddr.Mask[0] != 0xff || castaddr.Mask[1] != 0xff {
		fmt.Printf("\n\033[38;5;202m[Warning] Skipping \033[38;5;50m[\033[38;5;15m%s\033[38;5;50m] \033[38;5;196m(MASK Too large)", castaddr)
		return false
	}
	return true

}
