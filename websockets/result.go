package websockets

import "github.com/diogox/GoLauncher/api"

func ExtensionResultNew(title string, description string, icon string, onEnterAction api.Action, onAltEnterAction api.Action) *ExtensionResult {
	return &ExtensionResult{
		title:            title,
		description:      description,
		icon:             icon,
		onEnterAction:    onEnterAction,
		onAltEnterAction: onAltEnterAction,
	}
}

type ExtensionResult struct {
	title            string
	description      string
	icon             string
	onEnterAction    api.Action
	onAltEnterAction api.Action
}

func (er *ExtensionResult) Title() string {
	return er.title
}

func (er *ExtensionResult) Description() string {
	return er.description
}

func (er *ExtensionResult) IconPath() string {
	return er.icon
}

func (er *ExtensionResult) OnEnterAction() api.Action {
	return er.onEnterAction
}

func (er *ExtensionResult) OnAltEnterAction() api.Action {
	return er.onAltEnterAction
}
