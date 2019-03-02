package api

type Result interface {
	Title() string
	Description() string
	IconPath() string
	OnEnterAction() Action
	OnAltEnterAction() Action
}