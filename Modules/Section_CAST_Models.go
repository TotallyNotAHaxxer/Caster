package CastHunter

import "encoding/xml"

type DetailedResponseInfo struct {
	Bssid             string `json:"bssid"`
	BuildVersion      string `json:"build_version"`
	CastBuildRevision string `json:"cast_build_revision"`
	Connected         bool   `json:"connected"`
	EthernetConnected bool   `json:"ethernet_connected"`
	HasUpdate         bool   `json:"has_update"`
	HotspotBssid      string `json:"hotspot_bssid"`
	IPAddress         string `json:"ip_address"`
	Locale            string `json:"locale"`
	Location          struct {
		CountryCode string  `json:"country_code"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	} `json:"location"`
	MacAddress string `json:"mac_address"`
	Name       string `json:"name"`
	OptIn      struct {
		Crash    bool `json:"crash"`
		Opencast bool `json:"opencast"`
		Stats    bool `json:"stats"`
	} `json:"opt_in"`
	PublicKey    string `json:"public_key"`
	ReleaseTrack string `json:"release_track"`
	SetupState   int    `json:"setup_state"`
	SetupStats   struct {
		HistoricallySucceeded    bool `json:"historically_succeeded"`
		NumCheckConnectivity     int  `json:"num_check_connectivity"`
		NumConnectWifi           int  `json:"num_connect_wifi"`
		NumConnectedWifiNotSaved int  `json:"num_connected_wifi_not_saved"`
		NumInitialEurekaInfo     int  `json:"num_initial_eureka_info"`
		NumObtainIP              int  `json:"num_obtain_ip"`
	} `json:"setup_stats"`
	SsdpUdn       string  `json:"ssdp_udn"`
	Ssid          string  `json:"ssid"`
	TimeFormat    int     `json:"time_format"`
	TosAccepted   bool    `json:"tos_accepted"`
	UmaClientID   string  `json:"uma_client_id"`
	Uptime        float64 `json:"uptime"`
	Version       int     `json:"version"`
	WpaConfigured bool    `json:"wpa_configured"`
	WpaState      int     `json:"wpa_state"`
}

type InfoResponseUnDetailed struct {
	Bssid             string `json:"bssid"`
	BuildVersion      string `json:"build_version"`
	CastBuildRevision string `json:"cast_build_revision"`
	Connected         bool   `json:"connected"`
	EthernetConnected bool   `json:"ethernet_connected"`
	HasUpdate         bool   `json:"has_update"`
	HotspotBssid      string `json:"hotspot_bssid"`
	IPAddress         string `json:"ip_address"`
	Locale            string `json:"locale"`
	Location          struct {
		CountryCode string  `json:"country_code"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	} `json:"location"`
	MacAddress string `json:"mac_address"`
	Name       string `json:"name"`
	OptIn      struct {
		Crash    bool `json:"crash"`
		Opencast bool `json:"opencast"`
		Stats    bool `json:"stats"`
	} `json:"opt_in"`
	PublicKey    string `json:"public_key"`
	ReleaseTrack string `json:"release_track"`
	SetupState   int    `json:"setup_state"`
	SetupStats   struct {
		HistoricallySucceeded    bool `json:"historically_succeeded"`
		NumCheckConnectivity     int  `json:"num_check_connectivity"`
		NumConnectWifi           int  `json:"num_connect_wifi"`
		NumConnectedWifiNotSaved int  `json:"num_connected_wifi_not_saved"`
		NumInitialEurekaInfo     int  `json:"num_initial_eureka_info"`
		NumObtainIP              int  `json:"num_obtain_ip"`
	} `json:"setup_stats"`
	SsdpUdn       string  `json:"ssdp_udn"`
	Ssid          string  `json:"ssid"`
	TimeFormat    int     `json:"time_format"`
	TosAccepted   bool    `json:"tos_accepted"`
	UmaClientID   string  `json:"uma_client_id"`
	Uptime        float64 `json:"uptime"`
	Version       int     `json:"version"`
	WpaConfigured bool    `json:"wpa_configured"`
	WpaState      int     `json:"wpa_state"`
}

type SSDPXMLRESPONSEENTITY struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Link    []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Rel  string `xml:"rel,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"link"`
	Style []struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"style"`
	SpecVersion struct {
		Text  string `xml:",chardata"`
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"specVersion"`
	URLBase string `xml:"URLBase"`
	Device  struct {
		Text         string `xml:",chardata"`
		DeviceType   string `xml:"deviceType"`
		FriendlyName string `xml:"friendlyName"`
		Manufacturer string `xml:"manufacturer"`
		ModelName    string `xml:"modelName"`
		UDN          string `xml:"UDN"`
		IconList     struct {
			Text string `xml:",chardata"`
			Icon struct {
				Text     string `xml:",chardata"`
				Mimetype string `xml:"mimetype"`
				Width    string `xml:"width"`
				Height   string `xml:"height"`
				Depth    string `xml:"depth"`
				URL      string `xml:"url"`
			} `xml:"icon"`
		} `xml:"iconList"`
		ServiceList struct {
			Text    string `xml:",chardata"`
			Service struct {
				Text        string `xml:",chardata"`
				ServiceType string `xml:"serviceType"`
				ServiceId   string `xml:"serviceId"`
				ControlURL  string `xml:"controlURL"`
				EventSubURL string `xml:"eventSubURL"`
				SCPDURL     string `xml:"SCPDURL"`
			} `xml:"service"`
		} `xml:"serviceList"`
	} `xml:"device"`
}

type ATTROUTERDEVINFO struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Link    []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Rel  string `xml:"rel,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"link"`
	Style []struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"style"`
	SpecVersion struct {
		Text  string `xml:",chardata"`
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"specVersion"`
	Device struct {
		Text             string `xml:",chardata"`
		DeviceType       string `xml:"deviceType"`
		FriendlyName     string `xml:"friendlyName"`
		Manufacturer     string `xml:"manufacturer"`
		ManufacturerURL  string `xml:"manufacturerURL"`
		ModelDescription string `xml:"modelDescription"`
		ModelName        string `xml:"modelName"`
		ModelNumber      string `xml:"modelNumber"`
		ModelURL         string `xml:"modelURL"`
		SerialNumber     string `xml:"serialNumber"`
		UDN              string `xml:"UDN"`
		UPC              string `xml:"UPC"`
		ServiceList      struct {
			Text    string `xml:",chardata"`
			Service struct {
				Text        string `xml:",chardata"`
				ServiceType string `xml:"serviceType"`
				ServiceId   string `xml:"serviceId"`
				SCPDURL     string `xml:"SCPDURL"`
				ControlURL  string `xml:"controlURL"`
				EventSubURL string `xml:"eventSubURL"`
			} `xml:"service"`
		} `xml:"serviceList"`
		DeviceList struct {
			Text   string `xml:",chardata"`
			Device struct {
				Text             string `xml:",chardata"`
				DeviceType       string `xml:"deviceType"`
				FriendlyName     string `xml:"friendlyName"`
				Manufacturer     string `xml:"manufacturer"`
				ManufacturerURL  string `xml:"manufacturerURL"`
				ModelDescription string `xml:"modelDescription"`
				ModelName        string `xml:"modelName"`
				ModelNumber      string `xml:"modelNumber"`
				ModelURL         string `xml:"modelURL"`
				SerialNumber     string `xml:"serialNumber"`
				UDN              string `xml:"UDN"`
				UPC              string `xml:"UPC"`
				ServiceList      struct {
					Text    string `xml:",chardata"`
					Service struct {
						Text        string `xml:",chardata"`
						ServiceType string `xml:"serviceType"`
						ServiceId   string `xml:"serviceId"`
						SCPDURL     string `xml:"SCPDURL"`
						ControlURL  string `xml:"controlURL"`
						EventSubURL string `xml:"eventSubURL"`
					} `xml:"service"`
				} `xml:"serviceList"`
				DeviceList struct {
					Text   string `xml:",chardata"`
					Device struct {
						Text             string `xml:",chardata"`
						DeviceType       string `xml:"deviceType"`
						FriendlyName     string `xml:"friendlyName"`
						Manufacturer     string `xml:"manufacturer"`
						ManufacturerURL  string `xml:"manufacturerURL"`
						ModelDescription string `xml:"modelDescription"`
						ModelName        string `xml:"modelName"`
						ModelNumber      string `xml:"modelNumber"`
						ModelURL         string `xml:"modelURL"`
						SerialNumber     string `xml:"serialNumber"`
						UDN              string `xml:"UDN"`
						UPC              string `xml:"UPC"`
						ServiceList      struct {
							Text    string `xml:",chardata"`
							Service struct {
								Text        string `xml:",chardata"`
								ServiceType string `xml:"serviceType"`
								ServiceId   string `xml:"serviceId"`
								SCPDURL     string `xml:"SCPDURL"`
								ControlURL  string `xml:"controlURL"`
								EventSubURL string `xml:"eventSubURL"`
							} `xml:"service"`
						} `xml:"serviceList"`
					} `xml:"device"`
				} `xml:"deviceList"`
			} `xml:"device"`
		} `xml:"deviceList"`
		PresentationURL string `xml:"presentationURL"`
	} `xml:"device"`
}

type ATTROUTERSTRUCTURE struct {
	XMLName xml.Name `xml:"scpd"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Link    []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Rel  string `xml:"rel,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"link"`
	Style []struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"style"`
	SpecVersion struct {
		Text  string `xml:",chardata"`
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"specVersion"`
	ActionList struct {
		Text   string `xml:",chardata"`
		Action []struct {
			Text         string `xml:",chardata"`
			Name         string `xml:"name"`
			ArgumentList struct {
				Text     string `xml:",chardata"`
				Argument []struct {
					Text                 string `xml:",chardata"`
					Name                 string `xml:"name"`
					Direction            string `xml:"direction"`
					RelatedStateVariable string `xml:"relatedStateVariable"`
				} `xml:"argument"`
			} `xml:"argumentList"`
		} `xml:"action"`
	} `xml:"actionList"`
	ServiceStateTable struct {
		Text          string `xml:",chardata"`
		StateVariable []struct {
			Text             string `xml:",chardata"`
			SendEvents       string `xml:"sendEvents,attr"`
			Name             string `xml:"name"`
			DataType         string `xml:"dataType"`
			AllowedValueList struct {
				Text         string   `xml:",chardata"`
				AllowedValue []string `xml:"allowedValue"`
			} `xml:"allowedValueList"`
		} `xml:"stateVariable"`
	} `xml:"serviceStateTable"`
}
