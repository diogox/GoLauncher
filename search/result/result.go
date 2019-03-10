package result

import (
	"github.com/diogox/GoLauncher/api"
)

func NewSearchResult(title string, descr string, iconPath string, onEnterAction api.Action, onAltEnterAction api.Action) SearchResult {
	return SearchResult{

		Title_:       title,
		Description_: descr,
		IconPath_:    iconPath,
		OnEnterAction_: onEnterAction,
		OnAltEnterAction_: onAltEnterAction,
	}
}

type SearchResult struct {
	Title_            string
	Description_      string
	IconPath_         string
	OnEnterAction_    api.Action
	OnAltEnterAction_ api.Action
}

func (r SearchResult) Title() string {
	return r.Title_
}

func (r SearchResult) Description() string {
	return r.Description_
}

func (r SearchResult) IconPath() string {
	return r.IconPath_
}

func (r SearchResult) OnEnterAction() api.Action {
	return r.OnEnterAction_
}

func (r SearchResult) OnAltEnterAction() api.Action {
	return r.OnAltEnterAction_
}
