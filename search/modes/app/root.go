package app

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/modes"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/diogox/GoLauncher/search/util"
)

func NewAppSearchMode(db *api.DB, searchModes []modes.SearchMode) AppSearchMode {

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
		if len(results) > 8 { // TODO: Extract this to 'Number of results' preference or config file!
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
