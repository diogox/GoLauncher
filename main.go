package main

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/models"
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"github.com/diogox/GoLauncher/sqlite"
	"github.com/diogox/GoLauncher/websockets"
	"github.com/diogox/LinuxApps"
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

	// Start Extension Server
	go websockets.StartExtensions(&db)

	// Get All Apps
	apps := LinuxApps.GetApps()
	for _, app := range apps {
		appInfo := models.AppInfo{
			Exec:        app.ExecName,
			Name:        app.Name,
			Description: app.Description,
			IconName:    app.IconName,
		}
		_ = db.AddApp(appInfo)
	}

	// Watch for app changes
	onChange := func(app *LinuxApps.AppInfo) error {
		appInfo := models.AppInfo{
			Name:        app.Name,
			Description: app.Description,
			IconName:    app.IconName,
			Exec:        app.ExecName,
		}
		err := db.AddApp(appInfo)
		if err != nil {
			// Probably already exists. We just need to update it
			if err := db.UpdateApp(appInfo); err != nil {
				panic(err)
			}
		}
		return nil
	}
	onRemove := func() error {
		// Remove all apps and add them back. TODO: Could this be more efficient?
		if err := db.RemoveAllApps(); err != nil {
			panic(err)
		}

		apps := LinuxApps.GetApps()
		for _, app := range apps {
			appInfo := models.AppInfo{
				Exec: app.ExecName,
				Name: app.Name,
				Description: app.Description,
				IconName: app.IconName,
			}

			err := db.AddApp(appInfo)
			if err != nil {
				panic(err)
			}
		}

		return nil
	}
	appWatcher := LinuxApps.NewAppWatcher(onChange, onRemove)
	go appWatcher.Start()

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

	// Make main function wait
	wg.Wait()
}
