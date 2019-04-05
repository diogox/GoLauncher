package actions

import "github.com/diogox/GoLauncher/api"

var renderResultListInstance *RenderResultList

// Rendering results relies on the GUI framework.
// To keep the action platform-agnostic, we need to set it up before using it.
func SetupRenderResultList(renderResultsCallback func([]api.SearchResult) error) {
	renderResultListInstance = &RenderResultList {
		Type: api.RENDER_RESULT_LIST_ACTION,
		renderCallback: renderResultsCallback,
	}
}

func NewRenderResultList(resultList []api.SearchResult) RenderResultList {
	if renderResultListInstance == nil {
		panic("You need to setup this action before you can use it!")
	}

	newInstance := *renderResultListInstance
	newInstance.ResultList = resultList
	return newInstance
}

type RenderResultList struct {
	Type string
	ResultList []api.SearchResult `json:"result_list"`
	renderCallback func([]api.SearchResult) error
}

func (rrl RenderResultList) GetType() string {
	return rrl.Type
}

func (RenderResultList) KeepAppOpen() bool {
	return true
}

func (r RenderResultList) Run() error {
	return r.renderCallback(r.ResultList)
}

