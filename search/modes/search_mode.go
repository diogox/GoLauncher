package modes

import "github.com/diogox/GoLauncher/api"

type SearchMode interface {
	IsEnabled(input string) bool
	HandleInput(input string) api.Action
	DefaultItems(input string) []api.SearchResult
}
