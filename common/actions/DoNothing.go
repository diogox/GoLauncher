package actions

func NewDoNothingAction() DoNothing {
	return DoNothing{}
}

type DoNothing struct {}

func (DoNothing) keepAppOpen() bool {
	return true
}

func (DoNothing) run() {
	return
}
