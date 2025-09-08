package lsp

type PublishDiagnosticNotification struct {
	Notification
	Params PublishDiagnosticParams `json:"params"`
}
type PublishDiagnosticParams struct {
	Uri         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range   Range  `json:"range"`
	Source  string `json:"source"`
	Message string `json:"message"`
}

func NewDiagnostic(uri string, diag []Diagnostic) PublishDiagnosticNotification {
	return PublishDiagnosticNotification{
		Notification{
			RPC:    "2.0",
			Method: "textDocument/publishDiagnostics",
		},
		PublishDiagnosticParams{
			Uri:         uri,
			Diagnostics: diag,
		},
	}
}
