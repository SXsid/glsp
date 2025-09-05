package dummycompiler

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
