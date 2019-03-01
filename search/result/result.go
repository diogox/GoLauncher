package result

import (
	"github.com/diogox/GoLauncher/common"
)

func NewSearchResult(title string, descr string, iconPath string, onEnter common.Action, onAltEnter common.Action) SearchResult {
	return SearchResult{

		title:       title,
		description: descr,
		iconPath:    iconPath,
		onEnter: onEnter,
		onAltEnter: onAltEnter,
	}
}

type SearchResult struct {
	title       string
	description string
	iconPath    string
	onEnter common.Action
	onAltEnter common.Action
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

func (r SearchResult) OnEnter() common.Action {
	return r.onEnter
}

func (r SearchResult) OnAltEnter() common.Action {
	return r.onAltEnter
}
