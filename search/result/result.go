package result

import (
	"github.com/diogox/GoLauncher/common"
)

func NewSearchResult(title string, descr string, iconPath string, onEnterAction common.Action, onAltEnterAction common.Action) SearchResult {
	return SearchResult{

		title:       title,
		description: descr,
		iconPath:    iconPath,
		onEnterAction: onEnterAction,
		onAltEnterAction: onAltEnterAction,
	}
}

type SearchResult struct {
	title       string
	description string
	iconPath    string
	onEnterAction common.Action
	onAltEnterAction common.Action
}

func (r SearchResult) Title() string {
	return r.title
}

func (r SearchResult) Description() string {
	return r.description
}

func (r SearchResult) IconPath() string {
	return r.iconPath
}

func (r SearchResult) OnEnterAction() common.Action {
	return r.onEnterAction
}

func (r SearchResult) OnAltEnterAction() common.Action {
	return r.onAltEnterAction
}
