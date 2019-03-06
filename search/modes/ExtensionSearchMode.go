package modes

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/sqlite"
)

func NewExtensionSearchMode(db *api.DB) ExtensionSearchMode {

	sqlite.LoadDefaultShortcuts(db)

	return ExtensionSearchMode{
		db: db,
	}
}

type ExtensionSearchMode struct {
	db *api.DB
}

func (ExtensionSearchMode) IsEnabled(input string) bool {
	return true // TODO
}

func (ExtensionSearchMode) HandleInput(input string) api.Action {
	panic("implement me")
}

func (ExtensionSearchMode) DefaultItems(input string) []api.Result {
	panic("implement me")
}

