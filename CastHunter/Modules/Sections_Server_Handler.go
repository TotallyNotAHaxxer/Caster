package CastHunter

import (
	"net/http"
)

const (
	ServerStaticFP = "Modules/Static/"
)

var Paths = map[string]string{
	"/":                                "home.html",
	"/home":                            "home.html",
	"/server-index/new/indexes":        "home.html",
	"/server-index/new/indexes-errors": "docs.html",
	"/check":                           "check.html",
	"/enum":                            "enumerate.html",
	"/set":                             "set.html",
	"/view":                            "view.html",
	"/cast":                            "cast.html",
	"/roku":                            "roku.html",
	"/devs":                            "devs.html",
	"/hotmaps":                         "hotmaps.html",
	"/rpi":                             "rpis.html",
	"/adevs":                           "adevs.html",
	"/azdevs":                          "azdevs.html",
	"/ro":                              "ro.html",
}

var ErrorPaths = map[string]string{
	"parsererror": "formerror.html",
	"NoReqSup":    "RequestError.html",
}

func Process_Incoming_ReqStream(writer http.ResponseWriter, requeststream *http.Request) {
	requestpath := requeststream.URL.Path
	switch requeststream.Method {
	case "GET":
		newpath := ServerStaticFP + Paths[requestpath]
		http.ServeFile(writer, requeststream, newpath)
	case "POST":
		// create and log HTTP POST FORM VALUES
		if X := requeststream.ParseForm(); X != nil {
			newfp := ServerStaticFP + ErrorPaths["parsererror"]
			http.ServeFile(writer, requeststream, newfp) // log error form
		}
	default:
		newsrv := ServerStaticFP + ErrorPaths["NoReqSup"]
		http.ServeFile(writer, requeststream, newsrv)
	}

}
