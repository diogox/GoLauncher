package result

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/gtk3/result"
)

func NewSearchResult(opts SearchResultOptions) SearchResult {

	return SearchResult{
		title:            opts.Title,
		description:      opts.Description,
		iconPath:         opts.IconPath,
		isDefaultSelect:  opts.IsDefaultSelect,
		onEnterAction:    opts.OnEnterAction,
		onAltEnterAction: opts.OnAltEnterAction,
	}
}

type SearchResult struct {
	title            string
	description      string
	iconPath         string
	isDefaultSelect  bool
	onEnterAction    api.Action
	onAltEnterAction api.Action
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

func (r SearchResult) IsDefaultSelect() bool {
	return r.isDefaultSelect
}

func (r SearchResult) OnEnterAction() api.Action {
	return r.onEnterAction
}

func (r SearchResult) OnAltEnterAction() api.Action {
	return r.onAltEnterAction
}

func (r SearchResult) ToResultItem() api.ResultItem {
	opts := result.ResultItemOptions{
		Title:       r.title,
		Description: r.description,
		IconPath:    r.iconPath,
	}

	resultItem := result.NewResultItem(opts)

	return api.ResultItem(&resultItem)
}
