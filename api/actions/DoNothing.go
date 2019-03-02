package actions

func NewDoNothing() DoNothing {
	return DoNothing{}
}

type DoNothing struct {}

func (DoNothing) KeepAppOpen() bool {
	return true
}

func (DoNothing) Run() {
	return
}
