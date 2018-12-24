package scanner

// PortInfo represents port with open status, only open status will mean Open == true
type PortInfo struct {
	Port int
	Open bool
}
