package CastHunter

import "net/http"

const (
	ServerStaticFP = "Modules/Static/"
)

var Paths = map[string]string{
	"/":     "home.html",
	"/home": "home.html",
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
