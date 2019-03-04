package actions

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

func NewOpenUrl(url string) OpenUrl {
	return OpenUrl {
		Url: url,
	}
}

type OpenUrl struct {
	Url string
}

func (OpenUrl) KeepAppOpen() bool {
	return false
}

func (o OpenUrl) Run() {
	fmt.Println("Opening url: " + o.Url)
	err := open.Start(o.Url)
	if err != nil {
		panic(err)
	}
}

