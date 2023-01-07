package CastHunter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Trace_Intrist_IPAPCO(target string) {
	rows := [][]string{}
	cols := []string{
		"IP",
		"Network",
		"Version",
		"City",
		"Region",
		"Region Code",
		"Country",
		"Country Code",
		"Currency",
		"Postal Code",
		"ASN",
		"Organization",
		"Dial code",
		"UTC offset",
		"tdl",
	}
	body, x := http.Get(fmt.Sprintf("https://ipinfo.io/%s/json", target))
	if x != nil {
		fmt.Print("Error when making request -> ", x)
	} else {
		defer body.Body.Close()
		var T1 Ipapi_co
		bv, x := ioutil.ReadAll(body.Body)
		if x != nil {
			fmt.Println("Error reading HTTP response body: ", x)
		}
		json.Unmarshal(bv, &T1)
		rows = append(rows, []string{
			T1.IP,
			T1.Network,
			T1.Version,
			T1.City,
			T1.Region,
			T1.RegionCode,
			T1.Country,
			T1.CountryCode,
			T1.Currency,
			T1.Postal,
			T1.Asn,
			T1.Org,
			T1.CountryCallingCode,
			T1.UtcOffset,
			T1.CountryTld,
		})
	}
	if rows != nil {
		DrawTableSepColBased(rows, cols)
	}
}
