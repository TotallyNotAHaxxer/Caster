package CastHunter

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type UUIDINFO struct {
	XMLName     xml.Name `xml:"root"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	SpecVersion struct {
		Text  string `xml:",chardata"`
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"specVersion"`
	Device struct {
		Text         string `xml:",chardata"`
		DeviceType   string `xml:"deviceType"`
		FriendlyName string `xml:"friendlyName"`
		Manufacturer string `xml:"manufacturer"`
		ModelName    string `xml:"modelName"`
		UDN          string `xml:"UDN"`
		ServiceList  struct {
			Text    string `xml:",chardata"`
			Service struct {
				Text        string `xml:",chardata"`
				ServiceType string `xml:"serviceType"`
				ServiceId   string `xml:"serviceId"`
				SCPDURL     string `xml:"SCPDURL"`
				ControlURL  string `xml:"controlURL"`
				EventSubURL string `xml:"eventSubURL"`
			} `xml:"service"`
		} `xml:"serviceList"`
	} `xml:"device"`
}

var UUIDCAPS []string
var UUIDHOSTS []string

func RequestUUIDInformation(target string) {
	// this function will take existing UUID's and compare them to their hosts
	ur := FireTVPaths["uuidinf"]
	var uuid, host string
	if Id.UUID == nil {
		ErrorsPackets["uuidnil"]()
		return
	} else {
		for i := 0; i < len(Id.UUID); i++ {
			v := strings.Split(Id.UUID[i], "@")
			for _ = range v {
				// first should be host
				if v[0] == target && v[0] != "" {
					host = strings.TrimSpace(v[0])
					UUIDHOSTS = append(UUIDHOSTS, host)
				}
				if v[1] != "" {
					// should be uuid
					uuid = strings.TrimSpace(v[1])
					UUIDCAPS = append(UUIDCAPS, uuid)
				}
			}
		}
		newurl := fmt.Sprintf(ur, host, uuid)
		// make request and extract information
		res, x := http.Get(newurl)
		if x != nil {
			ErrorHandler[10]()
			fmt.Print(x)
			return
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			var uuidresp UUIDINFO
			bv, _ := ioutil.ReadAll(res.Body)
			xml.Unmarshal(bv, &uuidresp)
			var rows [][]string
			rows = append(rows,
				[]string{"Device Type", uuidresp.Device.DeviceType},
				[]string{"Device Name", uuidresp.Device.FriendlyName},
				[]string{"Device manufacturer", uuidresp.Device.Manufacturer},
				[]string{"Device Model", uuidresp.Device.ModelName},
				[]string{"Device UDN", uuidresp.Device.UDN},
				[]string{"Device ServiceType", uuidresp.Device.ServiceList.Service.ServiceType},
				[]string{"Device ServiceID", uuidresp.Device.ServiceList.Service.ServiceId},
				[]string{"Device ServiceSCPDURL", uuidresp.Device.ServiceList.Service.SCPDURL},
				[]string{"Device ControlURL", uuidresp.Device.ServiceList.Service.ControlURL},
				[]string{"Device EventSub URL", uuidresp.Device.ServiceList.Service.EventSubURL},
			)
			if rows != nil {
				DrawVerticle(rows)
			}
		} else {
			ErrorHandler[20]()
			return
		}
	}
}
