package extensions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/websockets"
	"strings"
)

func NewExtensionSearchMode() *ExtensionSearchMode {

	extensionsServer := websockets.GetExtensionsServer()

	return &ExtensionSearchMode{
		extensionsServer: extensionsServer,
	}
}

type ExtensionSearchMode struct {
	extensionsServer *websockets.ExtensionsServer
}

func (esm *ExtensionSearchMode) IsEnabled(input string) bool {
	keyword, _ := getKeywordArgs(input)
	if strings.Contains(input, *keyword + " ") {
		controller := esm.extensionsServer.GetControllerByKeyword(*keyword)
		if controller != nil {
			return true
		}
	}
	return false
}

func (esm *ExtensionSearchMode) HandleInput(input string) api.Action {
	// TODO
	results := make([]api.SearchResult, 0)
	keyword, args := getKeywordArgs(input)
	if strings.Contains(input, *keyword + " ") {
		controller := esm.extensionsServer.GetControllerByKeyword(*keyword)
		return controller.HandleInput(strings.Join(args, " "))
	}
	return actions.NewRenderResultList(results)
}

func (*ExtensionSearchMode) DefaultItems(input string) []api.SearchResult {
	return make([]api.SearchResult, 0)
}

// TODO: Create a 'query' type for this?
func getKeywordArgs(input string) (*string, []string) {
	components := strings.Split(input, " ")

	if len(components) > 0 {
		return &components[0], components[1:]
	}
	return nil, nil
}
