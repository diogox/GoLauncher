package actions

import "fmt"

func NewOpen(filepath string) Open {
	return Open{
		Filepath: filepath,
	}
}

type Open struct {
	Filepath string `json:"file_path"`
}

func (Open) KeepAppOpen() bool {
	return false
}

func (o Open) Run() {
	fmt.Println("Opening file: " + o.Filepath)
}

