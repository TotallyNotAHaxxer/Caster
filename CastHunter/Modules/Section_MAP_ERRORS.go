package CastHunter

import (
	"fmt"
)

var ErrorARPMAP = map[string]func(){
	"127": func() {
		fmt.Println("[Warning] Skipping address, was local")
	},
	"0xff": func() {
		fmt.Println("[Warning] Skipping address, was too large")
	},
}
