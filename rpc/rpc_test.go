package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EncodedMsg struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	msg := EncodedMsg{
		true,
	}
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	res := EncodeMessage(msg)
	assert.Equal(t, expected, res)

	incomingMsg := "Content-Length: 19\r\n\r\n{\"method\":\"update\"}"
	_, method, err := DecodeMessage([]byte(incomingMsg))
	assert.Equal(t, nil, err)
	assert.Equal(t, "update", method)
}
