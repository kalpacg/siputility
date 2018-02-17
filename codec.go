package siputility

import (
	"bytes"
	"errors"
)

// Represents a SIP message as defined in rfc 3261, ,section 7
type Message struct {
	Method  string
	Uri     string
	Version string
	Headers []byte
	Body    []byte
}

// Decodes SIP packet to a Message struct
//
func Decode(packet []byte) Message {
	m := Message{}

	reqLine, _, _, _ := getElements(packet)
	methodB, uriB, versionB, _ := getRequestLineElements(reqLine)

	m.Method = string(methodB)
	m.Uri = string(uriB)
	m.Version = string(versionB)
	m.Version = string(versionB)
	return m
}

// Returns every element thats is split by CRLF
//
func getElements(packet []byte) ([]byte, [][]byte, []byte, error) {
	elements := bytes.SplitAfter(packet, []byte("\r\n"))

	n := len(elements)

	if n >= 1 {
		return elements[0], elements[1 : n-2], elements[n-1], nil
	} else {
		return nil, nil, nil, errors.New("Decode error")
	}
}

// Get all elements in the request line.
// Request-Line MUST statisfy
// Request-Line  =  Method SP Request-URI SP SIP-Version CRLF
// in ABNF notation. See rfc 3261, section 7.1
//
func getRequestLineElements(binary []byte) ([]byte, []byte, []byte, error) {
	requestLine := bytes.SplitAfter(binary, []byte(" "))

	if len(requestLine) != 3 {
		return nil, nil, nil, errors.New("Decode error")
	} else {
		return requestLine[0], requestLine[1], requestLine[2], nil
	}
}
