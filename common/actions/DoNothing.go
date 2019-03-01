package actions

func NewDoNothingAction() DoNothing {
	return DoNothing{}
}

type DoNothing struct {}

func (DoNothing) KeepAppOpen() bool {
	return true
}

func (DoNothing) Run() {
	return
}
