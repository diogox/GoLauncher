package actions

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

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
	err := open.Start(o.url)
	if err != nil {
		panic(err)
	}
}

