package actions

func NewSetUserQueryAction(query string, setQueryCallback func()) SetUserQuery {
	return SetUserQuery{
		query: query,
		setQueryCallback: setQueryCallback,
	}
}

type SetUserQuery struct {
	query string
	setQueryCallback func()
}

func (SetUserQuery) KeepAppOpen() bool {
	return true
}

func (s SetUserQuery) Run() {
	s.setQueryCallback()
}

