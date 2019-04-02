package main

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/models"
	"github.com/diogox/GoLauncher/gtk3"
	"github.com/diogox/GoLauncher/search"
	"github.com/diogox/GoLauncher/search/result"
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
	preferences.BindCallback(api.PreferenceHotkey, func(arg interface{}) {
		hotkey, _ := arg.(string)
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
	go func() {
		err := appWatcher.Start()
		if err != nil {
			panic(err)
		}
	}()

	// Instantiate Search
	search := search.NewSearch(&db)

	// Instantiate Launcher
	launcher = gtk3.NewLauncher(&prefs)

	// Handle input function
	onInput := func(input string) {

		// TODO: Probably not thread-safe, use channels instead?
		actionResult := search.HandleInput(input)
		actionResult.Run()
	}
	onEmptyInput := func() {

		// Get preference
		showFrequentApps, err := preferences.GetPreference(api.PreferenceShowFrequentApps)
		if err != nil {
			panic(err)
		}

		// Show frequent apps or not
		if showFrequentApps == api.PreferenceFALSE {
			return
		}

		/* TODO: Should I use the preference number?
		// Get n of results to show by default
		nOfResultsStr, err := preferences.GetPreference(api.PreferenceNResultsToShow)
		if err != nil {
			panic(err)
		}

		nOfResults, err := strconv.Atoi(nOfResultsStr)
		*/
		nOfResults := 3
		frequentApps, err := db.GetMostFrequentApps(nOfResults)
		if err != nil {
			panic(err)
		}

		// Generate results
		// TODO: Implement ResultMaker interface (also TODO) to turn apps into results
		frequentAppsResults := make([]api.Result, 0)
		for _, app := range frequentApps {
			action := actions.NewLaunchApp(app.Exec, &db)
			r := result.NewSearchResult(app.Name, app.Description, app.IconName, false, action, action)
			frequentAppsResults = append(frequentAppsResults, api.Result(&r))
		}

		// Show results
		_, _ = glib.IdleAdd(launcher.ShowResults, frequentAppsResults)
	}

	// Set actions for when the input is changed
	launcher.HandleInput(onInput, onEmptyInput)

	// Start Launcher
	err := launcher.Start()
	if err != nil {
		// TODO: Log
		panic(err)
	}

	// Make main function wait
	wg.Wait()
}
