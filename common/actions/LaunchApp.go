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

func (LaunchApp) keepAppOpen() bool {
	return false
}

func (a *LaunchApp) run() {
	fmt.Println("Executing App: " + a.exec)
}

