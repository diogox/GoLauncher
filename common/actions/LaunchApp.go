package actions

import "fmt"

func NewLaunchAppAction(exec string) LaunchApp {
	return LaunchApp{
		exec: exec,
	}
}

type LaunchApp struct {
	exec string
}

func (LaunchApp) KeepAppOpen() bool {
	return false
}

func (a LaunchApp) Run() {
	fmt.Println("Executing App: " + a.exec)
}

