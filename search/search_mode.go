package search

import "github.com/diogox/GoLauncher/common"

type SearchMode interface {
	IsEnabled(input string) bool
	HandleInput(input string) common.Action
	//DefaultItems() []common.Result
}
