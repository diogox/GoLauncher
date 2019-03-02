package search

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/common/actions"
	"github.com/diogox/GoLauncher/search/modes/app"
	"github.com/diogox/GoLauncher/search/modes/calc"
	"github.com/diogox/GoLauncher/search/modes/file"
	"github.com/diogox/GoLauncher/sqlite"
)

func NewSearch(db *sqlite.LauncherDB) Search {

	// Define available search modes
	searchModes := []SearchMode {
		file.NewFileBrowserMode(),
		calc.NewCalcSearchMode(),
		app.NewAppSearchMode(db),
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

	return actions.NewDoNothing()
}
