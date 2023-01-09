package CastHunter

/*

Note: Amazon Fire Stick or any fire based devices from amazon are really hard to enumerate | Using urls like the http://%s:53917/zc?action=getInfo&version=2.7.1

to grab information about the device was how we would access serial information using UpNp however for some reastion the URL used does not exactly work around the same way it used to

so we now have to both craft and look for SSDP packets from the host of an amazon fire stick or amazon device. This is not really ethical or logical to do in a programatic sense

so this means instead of listening for addresses that may just be from a certain host which had an OUI for amazon we would just open a listener to listen for SSDP packets

we can also craft and send those same exact packets to get a response, this is a better way of doing things.

*/
// FireStick uses a different way of accessing system data and properties with UpNp
var StandardPortsFireTV = map[string]string{
	"9080": "glrpc SERVER",
	"8009": "Standard Port",
}

var FireTVPaths = map[string]string{
	"info":    "http://%s:53917/zc?action=getInfo&version=2.7.1",
	"uuidinf": "http://%s:60000/upnp/dev/%s/desc",
}
