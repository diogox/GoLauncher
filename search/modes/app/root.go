package app

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/common/actions"
	"github.com/diogox/GoLauncher/search/modes/app/finder"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/diogox/GoLauncher/sqlite"
)

func NewAppSearchMode(db *sqlite.LauncherDB, launcher *common.Launcher) AppSearchMode {
	// Get App Info
	finder.FindApps(db)

	return AppSearchMode{
		db: db,
		launcher: launcher,
	}
}

// TODO: Has to implement `SearchMode`
type AppSearchMode struct {
	db *sqlite.LauncherDB
	launcher *common.Launcher
}

func (AppSearchMode) IsEnabled(input string) bool {
	return true // TODO: Change this
}

func (asm AppSearchMode) HandleInput(input string) common.Action {

	results := make([]common.Result, 0)
	apps, err := asm.db.FindAppByName(input)
	if err != nil {
		panic(err)
	}

	for _, app := range apps {
		if len(results) > 5 {
			break
		}

		action := actions.NewLaunchAppAction(app.Name)
		r := result.NewSearchResult(app.Name, app.Description, app.IconName, action, action)
		results = append(results, r)
	}

	return actions.NewRenderResultListAction(results, (*asm.launcher).ShowResults)
}

