package main

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"github.com/diogox/GoLauncher/sqlite"
	"sync"
)

var wg sync.WaitGroup

func main() { 
	
	wg.Add(1)

	db := sqlite.NewLauncherDB()

	// Instantiate Launcher
	launcher := gtk3.NewLauncher()
	launcher.BindHotkey("<Ctrl>space")

	// Instantiate Search
	l := common.Launcher(launcher)
	search := search.NewSearch(&db, &l)

	// Handle input function
	launcher.HandleInput(func(input string) {

		// TODO: Probably not thread-safe, use channels instead
		actionResult := search.HandleInput(input)
		actionResult.Run()
	})

	// Start Launcher
	launcher.Start()

	// Make main function wait
	wg.Wait()
}
