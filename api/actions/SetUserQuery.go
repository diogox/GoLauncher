package actions

import "github.com/diogox/GoLauncher/api"

var setQueryInstance *SetUserQuery

// Copying text to the clipboard may rely on the GUI framework.
// To keep the action platform-agnostic, we need to set it up before using it.
func SetupSetUserQuery(setQueryCallback func(string)) {
	setQueryInstance = &SetUserQuery {
		Type: api.SET_USER_QUERY_ACTION,
		setQueryCallback: setQueryCallback,
	}
}

func NewSetUserQuery(query string) SetUserQuery {
	if setQueryInstance == nil {
		panic("You need to setup this action before you can use it!")
	}

	newInstance := *setQueryInstance
	newInstance.Query = query
	return newInstance
}

type SetUserQuery struct {
	Type string
	Query string
	setQueryCallback func(string)
}

func (suq SetUserQuery) GetType() string {
	return suq.Type
}

func (SetUserQuery) KeepAppOpen() bool {
	return true
}

func (s SetUserQuery) Run() {
	s.setQueryCallback(s.Query)
}

