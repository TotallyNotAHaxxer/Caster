package CastHunter

var Keys = map[string]string{
	"home": "http://%s:8060/keypress/home",
}

var PathsRoku = map[string]string{
	"browse":  "http://%s:8060/search/browse?keyword=%s",
	"devinfo": "http://%s:8060/query/device-info",
}
