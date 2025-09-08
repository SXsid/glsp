package lsp

type TextDocumentDefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type TextDocumentDefinitionReponse struct {
	Response
	Result Location `json:"result"`
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

func NewTextDocumentefinatinResoponse(id int, uri string, position Position) TextDocumentDefinitionReponse {
	return TextDocumentDefinitionReponse{
		Response{
			"2.0",
			&id,
		},
		Location{
			Uri: uri,
			Range: Range{
				Position{
					position.Line - 1,
					0,
				},
				Position{
					position.Line - 1,
					0,
				},
			},
		},
	}
}
