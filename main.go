package main

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"github.com/diogox/GoLauncher/sqlite"
	"github.com/diogox/GoLauncher/websockets"
	"github.com/gotk3/gotk3/glib"
	"sync"
)

var launcher *gtk3.Launcher
var wg sync.WaitGroup

// TODO: To allow for dynamically changing the Hotkey binding, for example, we need a dedicated "Preferences" object
//  that runs the appropriate functions after preference changes and commits them to the database.

func main() { 
	
	wg.Add(1)

	sqliteDB := sqlite.NewLauncherDB()
	db := api.DB(&sqliteDB)

	preferences := gtk3.PreferencesNew(&db)
	preferences.BindHotkeyCallBack(func(hotkey string) {
		_, _ = glib.IdleAdd(func() {
			launcher.BindHotkey(hotkey)
		})
	})
	prefs := api.Preferences(&preferences)
	_ = prefs.SetPreference(api.PreferenceHotkey, db.GetPreference(api.PreferenceHotkey))

	// TODO: Start extension server

	// Instantiate Search
	search := search.NewSearch(&db)

	// Instantiate Launcher
	launcher = gtk3.NewLauncher(&prefs)

	// Handle input function
	launcher.HandleInput(func(input string) {

		// TODO: Probably not thread-safe, use channels instead
		actionResult := search.HandleInput(input)
		actionResult.Run()
	})

	// Start Launcher
	launcher.Start()

	websockets.StartExtensionsServer()

	// Make main function wait
	wg.Wait()
}
