package search

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/modes"
	"github.com/diogox/GoLauncher/search/modes/app"
	"github.com/diogox/GoLauncher/search/modes/calc"
	"github.com/diogox/GoLauncher/search/modes/extensions"
	"github.com/diogox/GoLauncher/search/modes/file"
	"github.com/diogox/GoLauncher/search/modes/shortcut"
)

func NewSearch(db api.DB) Search {

	// Define available search modes
	searchModes := []modes.SearchMode {
		file.NewFileBrowserMode(),
		calc.NewCalcSearchMode(),
		shortcut.NewShortcutSearchMode(db),
		extensions.NewExtensionSearchMode(),
	}

	// Initialize AppSearchMode with the other modes
	searchModes = append(searchModes, app.NewAppSearchMode(db, searchModes))

	return Search{
		modes: searchModes,
	}
}

type Search struct {
	modes []modes.SearchMode
}

func (s *Search) HandleInput(input string) api.Action {

	for _, mode := range s.modes {
		if mode.IsEnabled(input) {
			return mode.HandleInput(input)
		}
	}

	return actions.NewDoNothing()
}
