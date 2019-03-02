package actions

import "fmt"

func NewOpen(filepath string) Open {
	return Open{
		filepath: filepath,
	}
}

type Open struct {
	filepath string
}

func (Open) KeepAppOpen() bool {
	return false
}

func (o Open) Run() {
	fmt.Println("Opening file: " + o.filepath)
}

