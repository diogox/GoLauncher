package app

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/search/modes/app/finder"
	"github.com/diogox/GoLauncher/search/result"
	"strings"
)

func NewAppSearchMode() AppSearchMode {
	// Get App Info
	apps := finder.FindApps()

	return AppSearchMode{
		apps: apps,
	}
}

// TODO: Has to implement `SearchMode`
type AppSearchMode struct {
	apps []finder.AppInfo
}

func (AppSearchMode) IsEnabled(input string) bool {
	return true // TODO: Change this
}

func (asm AppSearchMode) HandleInput(input string) []common.Result {

	results := make([]common.Result, 0)
	for _, app := range asm.apps {
		if len(results) > 5 {
			break
		}

		if strings.Contains(app.Name, input) {

			r := result.NewSearchResult(app.Name, app.Description, "/")
			results = append(results, r)
		}
	}

	return results
}

