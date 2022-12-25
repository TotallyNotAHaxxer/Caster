package CastHunter

import (
	"fmt"
)

// Typical error style  -
/*
	[1]| Error (code)
	[1]|_________
	[1]         | (Error type)
	[1]         |_____________
	[1]         | (Error message)
*/

func LoadMsg(code, errortype, msg string) string {
	var ErrorMsgTMPL string
	ErrorMsgTMPL += "\033[38;5;196m[1]| Error (%s)\n"
	ErrorMsgTMPL += "\033[38;5;56m[1]|_________\n"
	ErrorMsgTMPL += "\033[38;5;196m[1]         | (Error|%s)\n"
	ErrorMsgTMPL += "\033[38;5;56m[1]         |_____________\n"
	ErrorMsgTMPL += "\033[38;5;51m[1]         | (Message | -> %s)\n"
	ErrorMsgTMPL = fmt.Sprintf(ErrorMsgTMPL, code, errortype, msg)
	return ErrorMsgTMPL
}

var ErrorHandler = map[int]func(){

	10: func() {
		fmt.Println(LoadMsg("10", "Response Error", "Could not get a valid response from the server during GET request method, this may have been a connection issue, host issue etc"))
	},
	20: func() {
		fmt.Println(LoadMsg("20", "Response Error", "Response code from URL was not 200, this may be due to connection issue, host issue, Timeout issues and more"))
	},
	30: func() {
		fmt.Println(LoadMsg("30", "New Request Error", "Could not create a new request for POST using http.NewRequest and with current client and set headers"))
	},
	40: func() {
		fmt.Println(LoadMsg("40", "Fufill Request Client Error", "Could not create and execute the current request for POST method using current data"))
	},
	50: func() {
		fmt.Println(LoadMsg("50", "Commit to http request CLIENT", "Could not commit the current request because of a client error"))
	},
	60: func() {
		fmt.Println(LoadMsg("60", "Table input/output error", "Could not draw table due to there being no data within the response array"))
	},
	70: func() {
		fmt.Println(LoadMsg("70", "XML Input/Output Error", "Could not Marshal the response from the URL into an XML unmarshal response code"))
	},
	80: func() {
		fmt.Println(LoadMsg("80", "Settings and Configuration Error", "Could not find the required settings for this controller, use command `set target=targetIPv4` to set the address, where targetIPv4 is the IPv4 address of your target"))
	},
	90: func() {
		fmt.Println(LoadMsg("90", "Settings and Configuration Error", "Could not find the required settings for this controller, please ensure you have set an application ID with command `set appid=12` where 12 is the desired APPID"))
	},
	100: func() {},
	120: func() {},
	130: func() {},
	140: func() {},
	150: func() {},
	160: func() {},
	170: func() {},
}
