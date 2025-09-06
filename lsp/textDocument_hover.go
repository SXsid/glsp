package lsp

type TextDocumentHoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type TextDocumentHoverReponse struct {
	Response
	Result Hover `json:"result"`
}
type Hover struct {
	Contents string `json:"contents"`
}

// it repsonse with type definiton of the text in postion
func NewHoverReponse(Id int, content string) TextDocumentHoverReponse {
	return TextDocumentHoverReponse{
		Response{
			"2.0",
			&Id,
		},
		Hover{
			content,
		},
	}
}
