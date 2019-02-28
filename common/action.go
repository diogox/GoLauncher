package common

type Action interface {
	keepAppOpen() bool
	run()
}
