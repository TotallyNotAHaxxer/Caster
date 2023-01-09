package CastHunter

import (
	"crypto/tls"
	"net/http"
)

var GoogleHost_Enum []string

var (
	JD_Params = map[string]string{
		"params": "now",
	}
	JD_Paramsfr = map[string]string{
		"params": "now",
	}
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
)
