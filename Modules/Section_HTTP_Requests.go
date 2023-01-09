package CastHunter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Requests module

var (
	HTTP  = 8008
	HTTPs = 8443
)

//ADtq...
// something to look at https://chromecastbg.alexmeub.com/

var RequestHeadMap = map[string]string{
	"httpcontent": "application/json",
	"httpsauth":   "caster.assassin",
}

func MakeDelete(URL string, transport *http.Transport, JD map[string]string) {
	jv, _ := json.Marshal(JD)
	req, x := http.NewRequest("DELETE", URL, bytes.NewReader(jv))
	if x != nil {
		ErrorHandler[30]()
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: transport}
	resp, x := client.Do(req)
	if x != nil {
		ErrorHandler[40]()
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("[+] Success: Request fulfilled")
	} else {
		ErrorHandler[20]()
		return
	}
}

func MakePost(URL string, transport *http.Transport, JsonData map[string]string, istarget bool) {
	jsonValue, _ := json.Marshal(JsonData)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	if err != nil {
		ErrorHandler[30]()
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		ErrorHandler[40]()
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		ErrorHandler[20]()
	} else {
		if istarget {
			OK[100]()
		}
	}
}
