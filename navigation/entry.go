package navigation

import "github.com/diogox/GoLauncher/api"

type NavigationItem struct {
	SearchResult api.SearchResult
	ResulItem    api.ResultItem
}

func newNavigationItem(searchResult api.SearchResult) *NavigationItem {
	return &NavigationItem{
		SearchResult: searchResult,
		ResulItem:    searchResult.ToResultItem(),
	}
}
