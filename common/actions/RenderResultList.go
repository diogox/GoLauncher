package actions

import "github.com/diogox/GoLauncher/common"

func NewRenderResultListAction(resultList []*common.Result, renderResultsCallback func()) RenderResultList {
	return RenderResultList{
		resultList: resultList,
		renderCallback: renderResultsCallback,
	}
}

type RenderResultList struct {
	resultList []*common.Result
	renderCallback func()
}

func (RenderResultList) KeepAppOpen() bool {
	return true
}

func (r RenderResultList) Run() {
	r.renderCallback()
}

