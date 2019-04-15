package app

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/modes"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/diogox/GoLauncher/search/util"
	"strconv"
)

func NewAppSearchMode(db api.DB, searchModes []modes.SearchMode) AppSearchMode {

	return AppSearchMode{
		db:          db,
		searchModes: searchModes,
	}
}

type AppSearchMode struct {
	db          api.DB
	searchModes []modes.SearchMode
}

func (AppSearchMode) IsEnabled(input string) bool {
	return true
}

func (asm AppSearchMode) HandleInput(input string) api.Action {

	results := make([]api.SearchResult, 0)
	apps, _ := asm.db.GetAllApps()

	// Get exec of the result chosen the last time this input was searched
	execSelectedLastTime, _ := asm.db.GetSession(input)

	for _, app := range apps {
		isDefault := false

		if app.Name =="Spotify" {
			x := 1
			x++
		}

		// If the exec is the same, it's the same result
		if app.Exec == execSelectedLastTime {
			isDefault = true
		}

		launchAppOptions := actions.LaunchAppOptions{
			Db: asm.db,
			Exec: app.Exec,
			SearchTermUsed: input,
		}

		opts := result.SearchResultOptions{
			Title:            app.Name,
			Description:      app.Description,
			IconPath:         app.IconName,
			IsDefaultSelect:  isDefault,
			OnEnterAction:    actions.NewLaunchApp(launchAppOptions),
			OnAltEnterAction: actions.NewLaunchApp(launchAppOptions),
		}

		r := result.NewSearchResult(opts)
		results = append(results, r)
	}

	results = util.GetBestMatches(input, results, 50)

	// If there are no apps matching the search
	if len(results) == 0 {

		// Look for default items
		for _, searchMode := range asm.searchModes {
			for _, item := range searchMode.DefaultItems(input) {
				results = append(results, item)
			}
		}
		// Return default items here, to prevent them from getting their order changed by fuzzy search
		return actions.NewRenderResultList(results)
	}

	// Get maximum number of app results to show
	maxNOfApps, _ := strconv.Atoi(asm.db.GetPreference(api.PreferenceNAppResults))

	// If results exceed maximum number allowed, remove the ones less likely to be useful
	// (the ones at the end of the slice)
	if len(results) > maxNOfApps {
		results = results[:maxNOfApps-1]
	}

	return actions.NewRenderResultList(results)
}

func (AppSearchMode) DefaultItems(input string) []api.SearchResult {
	return nil
}
