package actions

import "github.com/diogox/GoLauncher/api"

func NewDoNothing() DoNothing {
	return DoNothing{
		Type: api.DO_NOTHING_ACTION,
	}
}

type DoNothing struct {
	Type string
}

func (dn DoNothing) GetType() string {
	return dn.Type
}

func (DoNothing) KeepAppOpen() bool {
	return true
}

func (DoNothing) Run() error {
	return nil
}
