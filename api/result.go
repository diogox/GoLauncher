package api

type Result interface {
	Title() string
	Description() string
	IconPath() string
	IsDefaultSelect() bool
	OnEnterAction() Action
	OnAltEnterAction() Action
}
