package CastHunter

import "encoding/xml"

// GDDEVDESC
type Root struct {
	XMLName     xml.Name `xml:"root"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
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

//L3SCPD

type Scpd struct {
	XMLName     xml.Name `xml:"scpd"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
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
				Argument struct {
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
		StateVariable struct {
			Text       string `xml:",chardata"`
			SendEvents string `xml:"sendEvents,attr"`
			Name       string `xml:"name"`
			DataType   string `xml:"dataType"`
		} `xml:"stateVariable"`
	} `xml:"serviceStateTable"`
}

// WANCI

type ScpdV2 struct {
	XMLName     xml.Name `xml:"scpd"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
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

// WANIP

type WANIPKEY struct {
	XMLName     xml.Name `xml:"scpd"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
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
			DefaultValue      string `xml:"defaultValue"`
			AllowedValueRange struct {
				Text    string `xml:",chardata"`
				Minimum string `xml:"minimum"`
				Maximum string `xml:"maximum"`
			} `xml:"allowedValueRange"`
		} `xml:"stateVariable"`
	} `xml:"serviceStateTable"`
}
