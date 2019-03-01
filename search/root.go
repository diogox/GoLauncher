package search

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/common/actions"
	"github.com/diogox/GoLauncher/search/modes/app"
	"github.com/diogox/GoLauncher/sqlite"
)

func NewSearch(db *sqlite.LauncherDB, launcher *common.Launcher) Search {

	// Define available search modes
	searchModes := []SearchMode {
		app.NewAppSearchMode(db, launcher),
	}

	return Search{
		modes: searchModes,
	}
}

type Search struct {
	modes []SearchMode
}

func (s *Search) HandleInput(input string) common.Action {

	for _, mode := range s.modes {
		if mode.IsEnabled(input) {
			return mode.HandleInput(input)
		}
	}

	return actions.NewDoNothingAction()
}
