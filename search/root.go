package search

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/search/modes/app"
)

func NewSearch() Search {

	searchModes := make([]SearchMode, 0)

	appMode := SearchMode(app.NewAppSearchMode())
	searchModes = append(searchModes, appMode)

	return Search{
		modes: searchModes,
	}
}

type Search struct {
	modes []SearchMode
}

func (s *Search) HandleInput(input string) []common.Result {

	for _, mode := range s.modes {
		if mode.IsEnabled(input) {
			return mode.HandleInput(input)
		}
	}

	return make([]common.Result, 0)
}
