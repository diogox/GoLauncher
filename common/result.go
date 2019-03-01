package common

type Result interface {
	Title() string
	Description() string
	IconPath() string
	OnEnter() Action
	OnAltEnter() Action
}