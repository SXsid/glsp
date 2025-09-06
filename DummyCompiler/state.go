package dummycompiler

import "fmt"

type State struct {
	Files map[string]string
}

func NewState() State {
	return State{map[string]string{}}
}

func (s *State) AddFile(fileName, data string) {
	s.Files[fileName] = data
}

func (s *State) UpdateFile(fileName, updatedData string) {
	s.Files[fileName] = updatedData
}

func (s *State) Hover(uri string) string {
	return fmt.Sprintf("File:%s ,count:%d ", s.Files[uri], len(s.Files[uri]))
}
