package CastHunter

import (
	"fmt"
	"net"
)

func VerifyConnPossible(ip string, port int) {
	host := net.JoinHostPort(ip, fmt.Sprint(port))
	conn, x := net.Dial("tcp", host)
	if x != nil {
		fmt.Println("[-] Error: Could not connect to host -> ", host, " on port ", port)
		return
	} else {
		fmt.Println("[+] Successfully connected to host -> ", host, " on port ", port)
		if port == HTTPs {
			fmt.Println("[+] Information: Using HTTPS port for connection -> ", HTTPs)
		} else if port == HTTP {
			fmt.Println("[+] Information: Using HTTP port for connection -> ", HTTP)
		}
	}
	defer conn.Close()

}
