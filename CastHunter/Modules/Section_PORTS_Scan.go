package CastHunter

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func VerifyPortStep(proto, host, service string, port int, resultChannel chan P_Results, Wait *sync.WaitGroup) {
	defer Wait.Done()
	ScanResults := P_Results{
		Port: port,
	}
	addr := net.JoinHostPort(host, fmt.Sprint(port))
	connect, x := net.DialTimeout(proto, addr, 1*time.Second)
	if x != nil {
		ScanResults.State = false
		resultChannel <- ScanResults
		return
	}
	defer connect.Close()
	ScanResults.State = true
	if service == "Google ChromeCast" {
		GoogleHosts = append(GoogleHosts, []string{
			host,
			"From PortScanner no MAC",
			"Google Chromecast (Based on port)",
		})
	}
	resultChannel <- ScanResults
	return
}

func GetOpenPorts(hostname string, ports Range) {
	scanned, x := ScanAddress(hostname, ports)
	if x != nil {
		fmt.Println("Error -. ", x)
	}
	Out(scanned)
}
