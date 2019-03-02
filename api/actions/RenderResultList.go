package actions

import "github.com/diogox/GoLauncher/api"

var renderResultListInstance *RenderResultList

// Rendering results relies on the GUI framework.
// To keep the action platform-agnostic, we need to set it up before using it.
func SetupRenderResultList(renderResultsCallback func([]api.Result)) {
	renderResultListInstance = &RenderResultList {
		renderCallback: renderResultsCallback,
	}
}

func NewRenderResultList(resultList []api.Result) RenderResultList {
	if renderResultListInstance == nil {
		panic("You need to setup this action before you can use it!")
	}

	newInstance := *renderResultListInstance
	newInstance.resultList = resultList
	return newInstance
}

type RenderResultList struct {
	resultList []api.Result
	renderCallback func([]api.Result)
}

func (RenderResultList) KeepAppOpen() bool {
	return true
}

func (r RenderResultList) Run() {
	r.renderCallback(r.resultList)
}

