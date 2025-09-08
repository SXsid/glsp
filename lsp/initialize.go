package lsp

type IntializeRequest struct {
	Request
	Params IntializeRequestParams `json:"params"`
}

type IntializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// there are more on specs site
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync   int  `json:"textDocumentSync"`
	HoverProvider      bool `json:"hoverProvider"`
	DefinitionProvider bool `json:"definitionProvider"`
	CodeActionProvider bool `json:"codeActionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response{
			"2.0",
			&id,
		},
		InitializeResult{
			ServerCapabilities{
				1,
				true,
				true,
				true,
			},
			ServerInfo{
				"glsp",
				"1.0.0",
			},
		},
	}
}
