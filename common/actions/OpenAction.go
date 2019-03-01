package actions

import "fmt"

func NewOpenAction(filepath string) OpenAction {
	return OpenAction{
		filepath: filepath,
	}
}

type OpenAction struct {
	filepath string
}

func (OpenAction) KeepAppOpen() bool {
	return false
}

func (o OpenAction) Run() {
	fmt.Println("Opening file: " + o.filepath)
}

