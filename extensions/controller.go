package extensions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/models"
)

func ControllerNew(extension *models.Extension) *Controller {
	return &Controller{
		Extension: extension,
	}
}

type Controller struct {
	Extension *models.Extension
}

func (Controller) HandleInput(args string) api.Action {
	results := make([]api.Result, 0)

	var action api.Action
	if args == "" {
		action = actions.NewCopyToClipboard("NO ARGS!")
	} else {
		action = actions.NewCopyToClipboard(args)
	}

	results = append(results, ExtensionResultNew("HackerNews", "See the top stories on HackerNews!", "google", action, action))
	return actions.NewRenderResultList(results)
}