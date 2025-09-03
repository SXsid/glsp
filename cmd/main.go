package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	dummycompiler "github.com/SXsid/glsp/DummyCompiler"
	"github.com/SXsid/glsp/lsp"
	"github.com/SXsid/glsp/rpc"
)

type Method string

const (
	Initialize          Method = "initialize"
	TextDocumentDidOpen Method = "textDocument/didOpen"
	Shutdown            Method = "shutdown"
)

func main() {
	state := dummycompiler.NewState()
	logger := getLogger()
	logger.Println("glsp is started")
	scanner := bufio.NewScanner(os.Stdin)
	// Split stream ,request by request
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		data := scanner.Bytes()
		body, method, err := rpc.DecodeMessage(data)
		if err != nil {
			log.Printf("Error occured:%s", err.Error())
		}

		handleMessage(logger, state, Method(method), body)
	}
}

func handleMessage(logger *log.Logger, state dummycompiler.State, method Method, body []byte) {
	logger.Printf("A request of %s type came...", method)
	switch method {
	case Initialize:
		var request lsp.IntializeRequest
		if err := json.Unmarshal(body, &request); err != nil {
			logger.Printf("error file parseing:%s\n", err.Error())
		}
		logger.Printf("Connected to: %s~%s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		response := rpc.EncodeMessage(msg)
		// reply
		writer := os.Stdout
		_, err := writer.Write([]byte(response))
		if err != nil {
			logger.Printf("error while responding :%s", err.Error())
		}
	case TextDocumentDidOpen:
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(body, &request); err != nil {
			logger.Printf("error file parseing:%s\n", err.Error())
		}
		state.AddFile(request.Params.TextDocumentItem.Uri, request.Params.TextDocumentItem.Text)
	}
}

func getLogger() *log.Logger {
	fileName := "/home/shekhar/Personal/glsp/log.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("error while creating the log file ")
	}
	return log.New(file, "[ glsp ]", log.Ldate|log.Ltime|log.Lshortfile)
}
