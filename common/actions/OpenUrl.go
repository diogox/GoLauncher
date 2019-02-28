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

func (OpenUrlAction) keepAppOpen() bool {
	return false
}

func (o *OpenUrlAction) run() {
	fmt.Println("Opening url: " + o.url)
}

