package CastHunter

import (
	"fmt"
)

type Messages func()

// Message format
func Format(code int, message string) string {
	var mes string
	mes += "\033[38;5;204m(\033[38;5;86mSuccess|C%s\033[38;5;204m) \033[38;5;208m-> %s"
	mes = fmt.Sprintf(mes, fmt.Sprint(code), message)
	return mes
}

var OK = map[int]Messages{
	98: func() {
		fmt.Println(Format(98, "Telling the device to go launch an application | Will wait on thread for 5 seconds"))

	},
	99: func() {
		fmt.Println(Format(99, "Enumeration / Function attempted to be run on host"))
	},
	100: func() {
		fmt.Println(Format(100, "Enumeration / Function ran successfully on target"))
	}, // Attacking / Enumerating hostname
	120: func() {}, // Enumeration was successfull
	130: func() {}, // Loading hosts or settings was successfull
}
