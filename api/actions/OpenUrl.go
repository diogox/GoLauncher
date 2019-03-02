package actions

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

func NewOpenUrl(url string) OpenUrl {
	return OpenUrl {
		url: url,
	}
}

type OpenUrl struct {
	url string
}

func (OpenUrl) KeepAppOpen() bool {
	return false
}

func (o OpenUrl) Run() {
	fmt.Println("Opening url: " + o.url)
	err := open.Start(o.url)
	if err != nil {
		panic(err)
	}
}

