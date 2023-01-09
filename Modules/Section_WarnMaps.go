package CastHunter

import "fmt"

var Warns = map[string]WarnMapCode{
	"possiblelag": func() {
		fmt.Println("[Warning] The host that was enumerated or sent the request did not give a verified response, as there is no 100% | true way to tell if the device turned off, however code was 200 which means the request should have been completed.")
	},
}
