package actions

import "github.com/diogox/GoLauncher/api"

func NewActionList(actions []api.Action) ActionList {
	return ActionList{
		actions: actions,
	}
}

type ActionList struct {
	actions []api.Action
}

func (al ActionList) KeepAppOpen() bool {
	if len(al.actions) == 0 {
		return true
	}

	// If any of the actions returns true, return true
	for _, action := range al.actions {
		if action.KeepAppOpen() == true {
			return true
		}
	}

	return false
}

func (al ActionList) Run() {
	for _, action := range al.actions {
		action.Run()
	}
}
