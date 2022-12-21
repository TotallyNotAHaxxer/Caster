package CastHunter

import (
	"fmt"
)

// Defines URL's, Sequences, structures

var (
	Information1         = "http://%s:%s/ssdp/device-desc.xml"       // Information endpoint of chromecast [1st url (Device Information)]
	Information2         = "http://%s:%s/setup/eureka_info"          // Information endpoint of chromecast [2nd url (Device Information)]
	ConfiguratedNetworks = "https://%s:%s/setup/configured_networks" // Information endpoint of chromecast [3rd url (Wifi names)]
	DeviceScan           = "https://%s:%s/setup/scan_wifi"           // Device function endpoint of chromecast [1st (tell device to scan for wifi)]
	DeviceScanResults    = "https://%s:%s/setup/scan_results"        // Device function endpoint of chromecast [2nd (tell device to output scan results)]
	DeviceForget         = "https://%s:%s/setup/forget_wifi"         // Device function endpoint of chromecast [3rd (tell device to forget wifiname)]
	DeviceRaname         = "https://%s:%s/setup/set_eureka_info"     // Device function endpoint of chromecast [4th (tell device to rename itself)]
	DeviceReboot         = "https://%s:%s/setup/reboot"              // Device function endpoint of chromecast [5th (tell device to reboot)]
	DeviceReset          = "https://%s:%s/setup/reboot"              // Device function endpoint of chromecast [6th (tell device to factory reset)]
)

var DevicePorts = map[string]string{
	"https": "8443",
	"http":  "8008",
}

var DeviceApps = map[string]func(string, string) string{
	"Youtube": func(ip, port string) string {
		return fmt.Sprintf("http://%s:%s/apps/YouTube", ip, port)
	},
	"Netflix": func(ip, port string) string {
		return fmt.Sprintf("http://%s:%s/apps/Netflix", ip, port)
	},
}

const (
	NULL = 0xff
	MVB  = 0
)

var (
	StartArp  bool
	ArpActive bool
)

var (
	MACS, IPS, OUIs         []string
	OutputRows_HostsANDmacs [][]string
	SleepInterval           int
	TargetMain              string
)
