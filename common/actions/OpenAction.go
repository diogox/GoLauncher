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

func (OpenAction) keepAppOpen() bool {
	return false
}

func (o *OpenAction) run() {
	fmt.Println("Opening file: " + o.filepath)
}

