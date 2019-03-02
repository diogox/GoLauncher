package actions

func NewHideWindow() *HideWindow {
	return &HideWindow{}
}

type HideWindow struct {}

func (HideWindow) KeepAppOpen() bool {
	return false
}

func (HideWindow) Run() {
	return
}

