package siputility

import (
	"bytes"
	"errors"
	"fmt"
)

// Represents a SIP message as defined in rfc 3261, ,section 7
type message struct {
	header byte
	body   byte
}

// Returns every element thats is split by CRLF
//
func getElements(packet []byte) ([][]byte, error) {
	elements := bytes.SplitAfter(packet, []byte("\r\n"))

	if len(elements) >= 1 {
		return elements, nil
	} else {
		return nil, errors.New("Decode error")
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
