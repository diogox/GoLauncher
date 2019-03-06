package actions

import (
	"fmt"
	"github.com/diogox/GoLauncher/api"
	"os/exec"
	"strings"
)

func NewLaunchApp(exec string) LaunchApp {
	return LaunchApp{
		Type: api.LAUNCH_APP_ACTION,
		Exec: exec,
	}
}

type LaunchApp struct {
	Type string
	Exec string
}

func (la LaunchApp) GetType() string {
	return la.Type
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

