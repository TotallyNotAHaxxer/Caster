package CastHunter

import (
	"fmt"
	"net/url"
	"strings"
)

// Parse URL's

func ExtractUUID(URL string, ssdp bool) string {
	l, x := url.Parse(URL)
	if x != nil {
		fmt.Println(x)
		return ""
	}
	if ssdp {
		port := l.Port()
		host := l.Host
		if port == "60000" {
			splitter := strings.Split(l.Path, "/")
			hostxUUID := host + "@" + splitter[3]
			return hostxUUID
			// Explanation:
			/*
				Within this brick of code since we are searching for a specific URL from SSDP we expect given UpNp for the URL to have 3 different fragments within the URL and nothing more

				this ofc is bad practice because one URL may look like a URL we want when it is not - TODO: Fix

				we use @ for table and formatting in terms of recongizing the hostname with the UUID, so we can say this IP address has this UUID for this url for this SSDP function and for this block
			*/
		}
	}
	return "portnotUPNPUUIDS----"
}
