package CastHunter

import "encoding/xml"

type ActiveApp struct {
	XMLName xml.Name `xml:"active-app"`
	Text    string   `xml:",chardata"`
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
	App []struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Type    string `xml:"type,attr"`
		Version string `xml:"version,attr"`
	} `xml:"app"`
}

type Apps struct {
	XMLName xml.Name `xml:"apps"`
	Text    string   `xml:",chardata"`
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
	App []struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Type    string `xml:"type,attr"`
		Version string `xml:"version,attr"`
	} `xml:"app"`
}

type DeviceInfo struct {
	XMLName xml.Name `xml:"device-info"`
	Text    string   `xml:",chardata"`
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
	Udn                      string `xml:"udn"`
	SerialNumber             string `xml:"serial-number"`
	DeviceID                 string `xml:"device-id"`
	AdvertisingID            string `xml:"advertising-id"`
	VendorName               string `xml:"vendor-name"`
	ModelName                string `xml:"model-name"`
	ModelNumber              string `xml:"model-number"`
	ModelRegion              string `xml:"model-region"`
	IsTv                     string `xml:"is-tv"`
	IsStick                  string `xml:"is-stick"`
	MobileHasLiveTv          string `xml:"mobile-has-live-tv"`
	UiResolution             string `xml:"ui-resolution"`
	SupportsEthernet         string `xml:"supports-ethernet"`
	WifiMac                  string `xml:"wifi-mac"`
	WifiDriver               string `xml:"wifi-driver"`
	HasWifiExtender          string `xml:"has-wifi-extender"`
	HasWifi5GSupport         string `xml:"has-wifi-5G-support"`
	CanUseWifiExtender       string `xml:"can-use-wifi-extender"`
	NetworkType              string `xml:"network-type"`
	NetworkName              string `xml:"network-name"`
	FriendlyDeviceName       string `xml:"friendly-device-name"`
	FriendlyModelName        string `xml:"friendly-model-name"`
	DefaultDeviceName        string `xml:"default-device-name"`
	UserDeviceName           string `xml:"user-device-name"`
	UserDeviceLocation       string `xml:"user-device-location"`
	BuildNumber              string `xml:"build-number"`
	SoftwareVersion          string `xml:"software-version"`
	SoftwareBuild            string `xml:"software-build"`
	SecureDevice             string `xml:"secure-device"`
	Language                 string `xml:"language"`
	Country                  string `xml:"country"`
	Locale                   string `xml:"locale"`
	TimeZoneAuto             string `xml:"time-zone-auto"`
	TimeZone                 string `xml:"time-zone"`
	TimeZoneName             string `xml:"time-zone-name"`
	TimeZoneTz               string `xml:"time-zone-tz"`
	TimeZoneOffset           string `xml:"time-zone-offset"`
	ClockFormat              string `xml:"clock-format"`
	Uptime                   string `xml:"uptime"`
	PowerMode                string `xml:"power-mode"`
	SupportsSuspend          string `xml:"supports-suspend"`
	SupportsFindRemote       string `xml:"supports-find-remote"`
	FindRemoteIsPossible     string `xml:"find-remote-is-possible"`
	SupportsAudioGuide       string `xml:"supports-audio-guide"`
	SupportsRva              string `xml:"supports-rva"`
	HasHandsFreeVoiceRemote  string `xml:"has-hands-free-voice-remote"`
	DeveloperEnabled         string `xml:"developer-enabled"`
	KeyedDeveloperID         string `xml:"keyed-developer-id"`
	SearchEnabled            string `xml:"search-enabled"`
	SearchChannelsEnabled    string `xml:"search-channels-enabled"`
	VoiceSearchEnabled       string `xml:"voice-search-enabled"`
	NotificationsEnabled     string `xml:"notifications-enabled"`
	NotificationsFirstUse    string `xml:"notifications-first-use"`
	SupportsPrivateListening string `xml:"supports-private-listening"`
	HeadphonesConnected      string `xml:"headphones-connected"`
	SupportsAudioSettings    string `xml:"supports-audio-settings"`
	SupportsEcsTextedit      string `xml:"supports-ecs-textedit"`
	SupportsEcsMicrophone    string `xml:"supports-ecs-microphone"`
	SupportsWakeOnWlan       string `xml:"supports-wake-on-wlan"`
	SupportsAirplay          string `xml:"supports-airplay"`
	HasPlayOnRoku            string `xml:"has-play-on-roku"`
	HasMobileScreensaver     string `xml:"has-mobile-screensaver"`
	SupportURL               string `xml:"support-url"`
	GrandcentralVersion      string `xml:"grandcentral-version"`
	TrcVersion               string `xml:"trc-version"`
	TrcChannelVersion        string `xml:"trc-channel-version"`
	DavinciVersion           string `xml:"davinci-version"`
	AvSyncCalibrationEnabled string `xml:"av-sync-calibration-enabled"`
}

type TvChannel struct {
	XMLName xml.Name `xml:"tv-channel"`
	Text    string   `xml:",chardata"`
	Channel struct {
		Text                  string `xml:",chardata"`
		Number                string `xml:"number"`
		Name                  string `xml:"name"`
		Type                  string `xml:"type"`
		UserHidden            string `xml:"user-hidden"`
		ActiveInput           string `xml:"active-input"`
		SignalState           string `xml:"signal-state"`
		SignalMode            string `xml:"signal-mode"`
		SignalQuality         string `xml:"signal-quality"`
		SignalStrength        string `xml:"signal-strength"`
		ProgramTitle          string `xml:"program-title"`
		ProgramDescription    string `xml:"program-description"`
		ProgramRatings        string `xml:"program-ratings"`
		ProgramAnalogAudio    string `xml:"program-analog-audio"`
		ProgramDigitalAudio   string `xml:"program-digital-audio"`
		ProgramAudioLanguages string `xml:"program-audio-languages"`
		ProgramAudioFormats   string `xml:"program-audio-formats"`
		ProgramAudioLanguage  string `xml:"program-audio-language"`
		ProgramAudioFormat    string `xml:"program-audio-format"`
		ProgramHasCc          string `xml:"program-has-cc"`
	} `xml:"channel"`
}
