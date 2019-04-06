package main

import (
	"fmt"
	"github.com/diogox/GoLauncher/cmd/launcher/app"
	"os"
)

func main() {
	cliApp := app.NewCliApp()
	if err := cliApp.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}