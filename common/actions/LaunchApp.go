package actions

import (
	"fmt"
	"os/exec"
	"strings"
)

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
	executable := strings.Split(a.exec, " ")
	cmd := exec.Command(executable[0], executable[1:]...)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

