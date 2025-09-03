// Package rpc is used for encoding and decoding the rpc messages
package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (body []byte, method string, err error) {
	header, content, isFound := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !isFound {
		return nil, "", errors.New("did not found seprator")
	}
	// find how long is the content
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return nil, "", err
	}
	// find what is the content
	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return nil, "", err
	}
	return content[:contentLength], baseMessage.Method, nil
}

type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// split incoming stream for sepration of reqeust
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, isFound := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !isFound {
		// no Error keep reading the data
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLength {
		// input stream haven't read full data which is coming so read keep reading
		return 0, nil, nil
	}

	// cause len(data ) rep whole buffer not the paritcular msg
	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
