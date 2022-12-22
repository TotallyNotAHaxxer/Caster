package CastHunter

import (
	"fmt"
	"strings"
)

type BoxFormat struct {
	TL string
	TR string
	BL string
	BR string
	HZ string
	VT string
}

func DrawUtilsBox(variable string) {
	BL := BoxFormat{
		TL: "┏",
		TR: "┓",
		BL: "┗",
		BR: "┛",
		HZ: "━",
		VT: "┃",
	}
	l := strings.Split(variable, "\n")
	var mlen int
	for _, lin := range l {
		if len(lin) > mlen {
			mlen = len(lin)
		}
	}
	fmt.Print("\033[38;5;93m" + BL.TL + strings.Repeat(BL.HZ, mlen) + BL.TR + "\n")
	for _, line := range l {
		fmt.Print(BL.VT + "\033[38;5;50m" + line + strings.Repeat(" ", mlen-len(line)) + "\033[38;5;93m" + BL.VT + "\n")
	}
	fmt.Print("\033[38;5;93m" + BL.BL + strings.Repeat(BL.HZ, mlen) + BL.BR + "\n")

}
