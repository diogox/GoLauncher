package actions

import "fmt"

func NewOpenUrlAction(url string) OpenUrlAction {
	return OpenUrlAction{
		url: url,
	}
}

type OpenUrlAction struct {
	url string
}

func (OpenUrlAction) KeepAppOpen() bool {
	return false
}

func (o OpenUrlAction) Run() {
	fmt.Println("Opening url: " + o.url)
}

