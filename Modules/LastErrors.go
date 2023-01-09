package CastHunter

import "fmt"

func ReturnMessage(code int, message string) string {
	var mes string
	mes += "\033[38;5;204m(\033[38;5;196mError|C%s\033[38;5;204m) \033[38;5;208m-> %s"
	mes = fmt.Sprintf(mes, fmt.Sprint(code), message)
	return mes
}

func ReturnDebug(code int, message string) string {
	var mes string
	mes += "\033[38;5;204m(\033[38;5;122mDEBUG|C%s\033[38;5;204m) \033[38;5;208m-> %s"
	mes = fmt.Sprintf(mes, fmt.Sprint(code), message)
	return mes
}

var ErrorStandard = map[int]func(){
	9901: func() {
		fmt.Println(ReturnMessage(9901, "Caster Capture: the SSDP module may not have been able to locate or find any UUID's at this moment | Is module running?"))
		// check module
		msg := " Is flag [--ssdp] true or enabled? "
		fmt.Println(ReturnDebug(9901, msg))
	},
	9920: func() {
		fmt.Println(ReturnMessage(9920, "Caster I/O: The array for the rows of the table was empty, maybe caster has not gotten all the data collected yet. Try again later "))
		fmt.Println(ReturnDebug(9920, "\tReason (#1): There may not be hosts discovered to view the ports, try waiting"))
		fmt.Println(ReturnDebug(9920, "\tReason (#2): Packets may not have been sent out yet, try waiting"))
		fmt.Println(ReturnDebug(9920, "\tReason (#3): The OUI's of the captured hosts do not match this field or input"))
		fmt.Println(ReturnDebug(9920, "\tReason (#4): The hosts or data required for this field may not exist on this network"))
	},
}
