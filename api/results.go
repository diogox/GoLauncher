package api

type SearchResult interface {
	Title() string
	Description() string
	IconPath() string
	IsDefaultSelect() bool
	OnEnterAction() Action
	OnAltEnterAction() Action
	ToResultItem() ResultItem
}

type ResultItem interface {
	Select()
	Unselect()
	SetPosition(position string)
	BindMouseHover(callback func())
	BindMouseClick(callback func())
	AccessInternals(callback func(args... interface{}))
}