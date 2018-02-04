package siputility

// Represents a SIP message as defined in rfc 3261, ,section 7
type message struct {
	header byte
	body   byte
}

// Decodes SIP binary
func Decode(binary []byte) *message {
	m := message{}
	m.header = binary[0]
	m.body = binary[1]
	return &m
}
