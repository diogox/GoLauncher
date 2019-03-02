package api

type Launcher interface {
	Start()
	Stop()
	BindHotkey(hotkey string)
	ToggleVisibility()
	ClearInput()
	HandleInput(func(input string))
	ShowResults([]Result)
}
