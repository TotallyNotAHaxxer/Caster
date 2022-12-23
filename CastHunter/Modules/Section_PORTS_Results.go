package CastHunter

import "fmt"

func Out(results Scan) {
	IPA := results.IP[len(results.IP)-1]
	fmt.Printf("\n\t\033[34m[\033[35m*\033[34m] Scan Results for   ├ %s (%s)\n", results.Hostname, IPA.String())
	Cols := []string{"Ports Open"}
	var Rows [][]string
	for _, v := range results.Results {
		if v.State {
			Rows = append(Rows, []string{fmt.Sprint(v.Port)})
		}
	}
	DrawTableSepColBased(Rows, Cols)

}

var common = map[int]string{
	8008: "Google ChromeCast",
	8443: "Google ChromeCast",
	8009: "Google ChromeCast",
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	161: "Simple Network Managment Protocol 	(SNMP)",
	162: "Simple Network Managment Protocol 	(SNMP)",
	389:   "LightWeight DIrectory Acess Protocol  (LDAP)",
	135:   "NetBIOS",
	49152: "CMS",
	65535: "CMS",
	49151: "Reserved",
	47808: "BACNET",
	44405: "Mu Online Connect Server",
	42806: "Discord",
	41797: "CSTP (Creston Secure Terminal Port)",
	41796: "CSCP (Creston Secure Control Port)",
	41795: "CTP  (Creston Terminal Port",
	41794: "CCP  (creston Control Port",
	41121: "Tentacal Server",
	40000: "SafetyNET p – a real-time Industrial Ethernet protocol",
	19812: "4D Database SQL Communication",
	19813: "4D database Client Server Communication",
	19814: "4D database DB4D Communication",
	19999: "DNP ( Distributed Network Protocal",
	19532: "SystemD-journal-gatewayd",
	19531: "SystemD-journal-remote",
	25565: "Minecraft (Java) Multiplayer Server",
	25575: "Minecraft (Java) Multiplayer Server RCON",
	23399: "Skype Server",
	22136: "FLIR Camera Resoruce Protocal",
	19302: "Google Talk/Video Communications",
	19295: "Google Talk/Video Communications",
	19294: "Google Talk/Video Communications",
	19133: "Minecraft Bedrock edition IPV6 multiplayer server",
	19132: "Minecraft Bedrock edition Multiplayer server",
	18333: "Bitcoing test network",
	16567: "BattleFeild 2 | mod ",
}
