package main

import (
	"github.com/diogox/raven/gtk3"
	"sync"
)

var wg sync.WaitGroup
var launcher gtk3.Launcher

func main() {

	wg.Add(1)

	// Instantiate launcher
	launcher = gtk3.NewLauncher()

	launcher.BindHotkey("<Ctrl>space")

	// Start Launcher
	launcher.Start()

	// Make main function wait
	wg.Wait()
}
