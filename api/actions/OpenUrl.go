package actions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/skratchdot/open-golang/open"
)

func NewOpenUrl(url string) OpenUrl {
	return OpenUrl {
		Type: api.OPEN_URL_ACTION,
		Url: url,
	}
}

type OpenUrl struct {
	Type string
	Url string
}

func (ou OpenUrl) GetType() string {
	return ou.Type
}

func (OpenUrl) KeepAppOpen() bool {
	return false
}

func (ou OpenUrl) Run() {
	err := open.Start(ou.Url)
	if err != nil {
		panic(err)
	}
}

