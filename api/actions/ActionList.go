package actions

import "github.com/diogox/GoLauncher/api"

func NewActionList(actions []api.Action) ActionList {
	return ActionList{
		Type: api.ACTION_LIST_ACTION,
		Actions: actions,
	}
}

type ActionList struct {
	Type string
	Actions []api.Action
}

func (al ActionList) GetType() string {
	return al.Type
}

func (al ActionList) KeepAppOpen() bool {
	if len(al.Actions) == 0 {
		return true
	}

	// If any of the actions returns true, return true
	for _, action := range al.Actions {
		if action.KeepAppOpen() == true {
			return true
		}
	}

	return false
}

func (al ActionList) Run() {
	for _, action := range al.Actions {
		action.Run()
	}
}

