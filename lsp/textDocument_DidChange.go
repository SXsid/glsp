package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}
type DidChangeTextDocumentParams struct {
	TextDocumentItem VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges   []TextDocumentContentChangeEvent `json:"contentChanges"`
}
