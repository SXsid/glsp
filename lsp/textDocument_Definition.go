package lsp

type TextDocumentDefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type TextDocumentDefinitionReponse struct {
	Response
	Result Location `json:"result"`
	Error  string   `json:"error"`
}
type Location struct {
	Uri   string `json:"uri"`
	Range Range  `json:"range"`
}

// this is repond with the file url and posiiton
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
