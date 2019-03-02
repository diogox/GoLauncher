package api

type Action interface {
	KeepAppOpen() bool
	Run()
}
