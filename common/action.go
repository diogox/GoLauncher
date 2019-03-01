package common

type Action interface {
	KeepAppOpen() bool
	Run()
}
