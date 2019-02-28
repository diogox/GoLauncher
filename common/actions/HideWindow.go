package actions

func NewHideWindowAction() HideWindow {
	return HideWindow{}
}

type HideWindow struct {}

func (HideWindow) keepAppOpen() bool {
	return false
}

func (HideWindow) run() {
	return
}

