package siputility

import (
	"bytes"
	"errors"
)

const (
	CRLF  = "\r\n"
	SP    = " "
	COLON = ":"
)

// Represents a SIP message as defined in rfc 3261, ,section 7
type Message struct {
	Method  string
	Uri     string
	Version string
	Headers []Header
	Body    []byte
}

// Represents a SIP header as defined in rfc 3261, ,section 7.3
type Header struct {
	Name  string
	Value string
}

// Decodes SIP packet to a Message struct
//
func Decode(packet []byte) Message {
	m := Message{}

	reqLine, headers, body, _ := getElements(packet)
	methodB, uriB, versionB, _ := getRequestLineElements(reqLine)
	m.Method = string(methodB)
	m.Uri = string(uriB)
	m.Version = string(versionB)
	m.Headers = getHeaders(headers)
	m.Body = body
	return m
}

// Encode a Message struct to SIP packet
//
func Encode(method, uri, version string, headers []Header, body []byte) []byte {

	methodB := []byte(method)
	uriB := []byte(uri)
	versionB := []byte(version)

	var headersB []byte

	for _, h := range headers {
		headersB = append(headersB, h.Name...)
		headersB = append(headersB, COLON...)
		headersB = append(headersB, h.Value...)
		headersB = append(headersB, CRLF...)
	}

	stream := [][]byte{methodB, []byte(SP), uriB, []byte(SP), versionB, []byte(CRLF),
		headersB, []byte(CRLF), body}

	return Concat(stream)
}

// Concatenates array of byte slices into one byte slice
func Concat(elements [][]byte) []byte {

	var stream []byte

	for _, element := range elements {
		for _, b := range element {
			stream = append(stream, b)
		}
	}

	return stream
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

// ABNF notation for SIP header is as follows. See rfc 3261, section 7.3
// header  =  "header-name" HCOLON header-value *(COMMA header-value)
// field-name: field-value
// field-name: field-value *(;parameter-name=parameter-value)
//
func getHeaders(headersB [][]byte) []Header {
	headers := []Header{}

	for _, hB := range headersB {
		elems := bytes.SplitN(hB, []byte(":"), 2)

		h := Header{Name: string(elems[0]), Value: string(elems[1])}
		headers = append(headers, h)
	}

	return headers
}

func GetHeaderValue(headers []Header, name string) (string, error) {
	for _, h := range headers {
		if h.Name == name {
			return h.Value, nil
		}
	}
	return "", errors.New("Header value not found")
}
