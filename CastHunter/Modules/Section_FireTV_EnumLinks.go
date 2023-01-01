package CastHunter

// FireStick uses a different way of accessing system data and properties with UpNp
var StandardPortsFireTV = map[string]string{
	"9080": "glrpc SERVER",
	"8009": "Standard Port",
}

var FireTVPaths = map[string]string{
	"info": "http://%s:53917/zc?action=getInfo&version=2.7.1",
}
