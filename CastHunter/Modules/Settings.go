package CastHunter

type Flags struct {
	Arp          bool
	SendOutEvery int
	OneInterface bool
	Server       bool // launch a local HTTP server
	Trace        bool // Traceback logging for panic
	Error        bool // Error logging for panic
}
