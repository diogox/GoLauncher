package main

import (
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"sync"
)

var wg sync.WaitGroup
var launcher gtk3.Launcher

func main() { 
	
	wg.Add(1)

	// Instantiate Search
	search := search.NewSearch()

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
