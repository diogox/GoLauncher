package main

type Launcher interface {
	Start()
	Stop()
	BindHotkey(hotkey string)
	ToggleVisibility()
	ClearInput()
	ShowResults([]Result)
}
