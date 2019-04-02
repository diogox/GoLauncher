package actions

import "github.com/diogox/GoLauncher/api"

func NewHideWindow() *HideWindow {
	return &HideWindow{
		Type: api.HIDE_WINDOW_ACTION,
	}
}

type HideWindow struct {
	Type string
}

func (hw HideWindow) GetType() string {
	return hw.Type
}

func (HideWindow) KeepAppOpen() bool {
	return false
}

func (HideWindow) Run() error {
	return nil
}

