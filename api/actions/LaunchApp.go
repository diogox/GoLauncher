package actions

import (
	"fmt"
	"os/exec"
	"strings"
)

func NewLaunchApp(exec string) LaunchApp {
	return LaunchApp{
		Exec: exec,
	}
}

type LaunchApp struct {
	Exec string
}

func (LaunchApp) KeepAppOpen() bool {
	return false
}

func (a LaunchApp) Run() {
	fmt.Println("Executing App: " + a.Exec)
	executable := strings.Split(a.Exec, " ")
	cmd := exec.Command(executable[0], executable[1:]...)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	_ = cmd.Process.Release()
}

