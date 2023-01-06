package CastHunter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func GrabSig(filename string) {
	f, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println(x)
		return
	}
	if x != nil {
		fmt.Println(x)
		return
	}
	var found bool
	for _, sig := range MassSign {
		if strings.HasSuffix(filename, sig.SuffixFile) || bytes.Contains(f, []byte(sig.Sign)) {
			var MB string
			if sig.FileFormat == "Apple Binary Property List (PLIST)" {
				MB = "62 70 6c 69 73 74 30 30"
			}
			fmt.Printf("\n\033[38;5;50m[\033[38;5;56mInformation\033[38;5;50m] Forensics module found file  \n \t -- NAME          | (\033[38;5;56m%v\033[38;5;50m) \n \t -- ASCII Sign    | (\033[38;5;56m%s\033[38;5;50m) \n \t -- Byte          | (\033[38;5;56m%v\033[38;5;50m) \n \t -- Magic Header  | (\033[38;5;56m%v\033[38;5;50m)", sig.FileFormat, sig.Sign, []byte(sig.Sign), MB)
			found = true
		}
		if found {
			break
		}
	}
}
