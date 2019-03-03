package main

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"github.com/diogox/GoLauncher/sqlite"
	"sync"
)

var wg sync.WaitGroup

// TODO: To allow for dynamically changing the Hotkey binding, for example, we need a dedicated "Preferences" object
//  that runs the appropriate functions after preference changes and commits them to the database.

func main() { 
	
	wg.Add(1)

	sqliteDB := sqlite.NewLauncherDB()
	db := api.DB(&sqliteDB)

	// Instantiate Search
	search := search.NewSearch(&db)

	// Instantiate Launcher
	launcher := gtk3.NewLauncher(&db)

	hotkey := db.GetPreference(api.PreferenceHotkey)
	launcher.BindHotkey(hotkey)

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
