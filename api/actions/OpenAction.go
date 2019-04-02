package actions

import (
	"fmt"
	"github.com/diogox/GoLauncher/api"
)

func NewOpen(filepath string) Open {
	return Open{
		Type: api.OPEN_ACTION,
		Filepath: filepath,
	}
}

type Open struct {
	Type string
	Filepath string `json:"file_path"`
}

func (o Open) GetType() string {
	return o.Type
}

func (Open) KeepAppOpen() bool {
	return false
}

func (o Open) Run() error {
	fmt.Println("Opening file: " + o.Filepath)

	// TODO

	return nil
}

