package actions

import (
	"github.com/diogox/GoLauncher/api"
	"os/exec"
	"strings"
)

func NewLaunchApp(exec string, db *api.DB) LaunchApp {
	return LaunchApp{
		Type: api.LAUNCH_APP_ACTION,
		Exec: exec,
		db:   db,
	}
}

type LaunchApp struct {
	Type string
	Exec string
	db   *api.DB
}

func (la LaunchApp) GetType() string {
	return la.Type
}

func (LaunchApp) KeepAppOpen() bool {
	return false
}

func (a LaunchApp) Run() {
	// Increment access counter
	err := (*a.db).IncrementAppAccessCounter(a.Exec)
	if err != nil {
		panic(err)
	}

	executable := strings.Split(a.Exec, " ")
	cmd := exec.Command(executable[0], executable[1:]...)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	_ = cmd.Process.Release()
}
