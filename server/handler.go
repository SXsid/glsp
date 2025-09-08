package server

import (
	"encoding/json"

	"github.com/SXsid/glsp/lsp"
	"github.com/SXsid/glsp/rpc"
)

func (s *Server) handleMessage(method Method, body []byte) {
	s.logger.Printf("A request of %s type came...", method)
	switch method {

	case Initialize:
		var request lsp.IntializeRequest
		if err := json.Unmarshal(body, &request); err != nil {
			s.logger.Printf("error file parseing:%s\n", err.Error())
		}
		s.logger.Printf("Connected to: %s~%s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		s.writeResponse(msg)
	case TextDocumentDidOpen:
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(body, &request); err != nil {
			s.logger.Printf("error file parseing:%s\n", err.Error())
		}
		s.logger.Printf("opend file :%s", request.Params.TextDocumentItem.Uri)
		s.State.AddFile(request.Params.TextDocumentItem.Uri, request.Params.TextDocumentItem.Text)

	case TextDidChange:
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(body, &request); err != nil {
			s.logger.Printf("error file parseing:%s\n", err.Error())
		}
		s.logger.Printf("changed file :%s", request.Params.TextDocumentItem.Uri)
		for _, change := range request.Params.ContentChanges {
			diagnostics := s.State.UpdateFile(request.Params.TextDocumentItem.Uri, change.NewData)
			msg := lsp.NewDiagnostic(request.Params.TextDocumentItem.Uri, diagnostics)
			s.writeResponse(msg)
		}

	case TextHover:

		var request lsp.TextDocumentHoverRequest
		if err := json.Unmarshal(body, &request); err != nil {
			s.logger.Printf("error file parseing:%s\n", err.Error())
		}
		content := s.State.Hover(request.Params.TextDocument.Uri)
		msg := lsp.NewHoverReponse(request.ID, content)
		s.writeResponse(msg)
	case TextDefinition:

		var request lsp.TextDocumentDefinitionRequest
		if err := json.Unmarshal(body, &request); err != nil {
			s.logger.Printf("error file parseing:%s\n", err.Error())
		}
		msg := lsp.NewTextDocumentefinatinResoponse(request.ID, request.Params.TextDocument.Uri, request.Params.Position)
		s.writeResponse(msg)
	}
}

func (s *Server) writeResponse(msg any) {
	response := rpc.EncodeMessage(msg)
	_, err := s.out.Write([]byte(response))
	if err != nil {
		s.logger.Printf("error while responding :%s", err.Error())
	}
}
