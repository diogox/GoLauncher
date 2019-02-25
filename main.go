package main

import (
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"github.com/diogox/GoLauncher/sqlite"
	"sync"
)

var wg sync.WaitGroup
var launcher gtk3.Launcher

func main() { 
	
	wg.Add(1)

	db := sqlite.NewLauncherDB()

	// Instantiate Search
	search := search.NewSearch(&db)

	// Instantiate Launcher
	launcher = gtk3.NewLauncher()
	launcher.BindHotkey("<Ctrl>space")
	launcher.HandleInput(func(input string) {

		// TODO: Probably not thread-safe, use channels instead
		searchResults := search.HandleInput(input)
		launcher.ShowResults(searchResults)
	})

	// Start Launcher
	launcher.Start()

	// Make main function wait
	wg.Wait()
}
