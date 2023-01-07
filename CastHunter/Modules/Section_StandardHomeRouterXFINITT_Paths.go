package CastHunter

// typical port is 49152
var RoutesARRIS = map[string]string{
	"WANIP":     "http://%s:49152/WANIPConnectionServiceSCPD.xml",
	"WANCI":     "http://%s:49152/WANCommonInterfaceConfigSCPD.xml",
	"L3SCPD":    "http://%s:49152/Layer3ForwardingSCPD.xml",
	"GDDEVDESC": "http://%s:49152/IGDdevicedesc_brlan0.xml",
}
