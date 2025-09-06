package server

import (
	"bufio"
	"io"
	"log"

	dummycompiler "github.com/SXsid/glsp/DummyCompiler"
	"github.com/SXsid/glsp/rpc"
)

type Method string

const (
	Initialize          Method = "initialize"
	TextDocumentDidOpen Method = "textDocument/didOpen"
	Shutdown            Method = "shutdown"
	TextDidChange       Method = "textDocument/didChange"
	TextHover           Method = "textDocument/hover"
	TextDefinition      Method = "textDocument/definition"
)

type Server struct {
	State  dummycompiler.State
	logger *log.Logger
	// to mock test
	in  io.Reader
	out io.Writer
}

func NewServer(State dummycompiler.State, logger *log.Logger, in io.Reader, out io.Writer) *Server {
	return &Server{
		State,
		logger,
		in,
		out,
	}
}

func (s *Server) Start() error {
	scanner := bufio.NewScanner(s.in)
	// INFO:find content length to split request form pool of request
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		data := scanner.Bytes()
		body, method, err := rpc.DecodeMessage(data)
		if err != nil {
			s.logger.Printf("decode error:%s", err)
			continue
		}
		// INFO:decode the body based on the baseMethod
		s.handleMessage(Method(method), body)
	}
	return scanner.Err()
}
