package app

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/models"
	"github.com/diogox/GoLauncher/search/modes"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/diogox/GoLauncher/search/util"
	"github.com/diogox/LinuxApps"
)

func NewAppSearchMode(db *api.DB, searchModes []modes.SearchMode) AppSearchMode {
	// Get App Info
	apps := LinuxApps.GetApps()
	for _, app := range apps {
		appInfo := models.AppInfo{
			Exec:        app.ExecName,
			Name:        app.Name,
			Description: app.Description,
			IconName:    app.IconName,
		}
		_ = (*db).AddApp(appInfo)
	}

	return AppSearchMode{
		db:          db,
		searchModes: searchModes,
	}
}

type AppSearchMode struct {
	db          *api.DB
	searchModes []modes.SearchMode
}

func (AppSearchMode) IsEnabled(input string) bool {
	return true
}

func (asm AppSearchMode) HandleInput(input string) api.Action {

	results := make([]api.Result, 0)
	apps, _ := (*asm.db).FindAppByName(input)

	for _, app := range apps {
		if len(results) > 5 {
			break
		}

		action := actions.NewLaunchApp(app.Exec)
		r := result.NewSearchResult(app.Name, app.Description, app.IconName, action, action)
		results = append(results, r)
	}

	if len(results) == 0 {
		for _, searchMode := range asm.searchModes {
			for _, item := range searchMode.DefaultItems(input) {
				results = append(results, item)
			}
		}
	}

	return actions.NewRenderResultList(util.GetBestMatches(input, results))
}

func (AppSearchMode) DefaultItems(input string) []api.Result {
	return nil
}
