package app

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/modes/app/finder"
	"github.com/diogox/GoLauncher/search/result"
)

func NewAppSearchMode(db *api.DB) AppSearchMode {
	// Get App Info
	finder.FindApps(db)

	return AppSearchMode{
		db: db,
	}
}

// TODO: Has to implement `SearchMode`
type AppSearchMode struct {
	db *api.DB
}

func (AppSearchMode) IsEnabled(input string) bool {
	return true // TODO: Change this
}

func (asm AppSearchMode) HandleInput(input string) api.Action {

	results := make([]api.Result, 0)
	apps, err := (*asm.db).FindAppByName(input)
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		if len(results) > 5 {
			break
		}

		action := actions.NewLaunchApp(app.Exec)
		r := result.NewSearchResult(app.Name, app.Description, app.IconName, action, action)
		results = append(results, r)
	}

	return actions.NewRenderResultList(results)
}

