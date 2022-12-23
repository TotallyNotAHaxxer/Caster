package CastHunter

import (
	"fmt"
	"net"
	"sync"
)

func ScanAddress(hostname string, rang Range) (Scan, error) {
	var (
		res     []P_Results
		scanned Scan
		Wait    sync.WaitGroup
	)
	resultChannel := make(chan P_Results, rang.End-rang.Start)
	addr, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("[-] Looks like there was an error looking up the address")
		return scanned, err
	}
	var added int
	for i := rang.Start; i <= rang.End; i++ {
		if ser, ok := common[i]; ok {
			Wait.Add(1)
			added++
			go VerifyPortStep("tcp", hostname, ser, i, resultChannel, &Wait)
		}
	}
	Wait.Wait()
	close(resultChannel)
	for result := range resultChannel {
		res = append(res, result)
	}
	scanned = Scan{
		Hostname: hostname,
		IP:       addr,
		Results:  res,
	}
	return scanned, nil
}
