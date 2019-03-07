package extensions

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/extensions"
	"strings"
)

func NewExtensionSearchMode(db *api.DB) *ExtensionSearchMode {

	controllers := make([]*extensions.Controller, 0)

	exts, _ := (*db).GetAllExtensions()
	for _, extension := range exts {
		controller := extensions.ControllerNew(&extension)
		controllers = append(controllers, controller)
	}

	return &ExtensionSearchMode{
		controllers: controllers,
	}
}

type ExtensionSearchMode struct {
	controllers []*extensions.Controller
}

func (esm *ExtensionSearchMode) IsEnabled(input string) bool {
	keyword, _ := getKeywordArgs(input)
	if strings.Contains(input, *keyword + " ") {
		_, err := esm.getControllerByKeyword(*keyword)
		if err == nil {
			return true
		}
	}
	return false
}

func (esm *ExtensionSearchMode) HandleInput(input string) api.Action {
	// TODO
	results := make([]api.Result, 0)
	keyword, args := getKeywordArgs(input)
	if strings.Contains(input, *keyword + " ") {
		controller, _ := esm.getControllerByKeyword(*keyword)
		return controller.HandleInput(strings.Join(args, " "))
	}
	return actions.NewRenderResultList(results)
}

func (*ExtensionSearchMode) DefaultItems(input string) []api.Result {
	return make([]api.Result, 0)
}

func (esm *ExtensionSearchMode) getControllerByKeyword(keyword string) (*extensions.Controller, error) {

	for _, controller := range esm.controllers {
		if controller.Extension.Keyword == keyword {
			return controller, nil
		}
	}

	return nil, errors.New("controller not found")
}

// TODO: Create a 'query' type for this?
func getKeywordArgs(input string) (*string, []string) {
	components := strings.Split(input, " ")

	if len(components) > 0 {
		return &components[0], components[1:]
	}
	return nil, nil
}
