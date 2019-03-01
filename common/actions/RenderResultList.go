package actions

import "github.com/diogox/GoLauncher/common"

func NewRenderResultListAction(resultList []common.Result, renderResultsCallback func([]common.Result)) RenderResultList {
	return RenderResultList{
		resultList: resultList,
		renderCallback: renderResultsCallback,
	}
}

type RenderResultList struct {
	resultList []common.Result
	renderCallback func([]common.Result)
}

func (RenderResultList) KeepAppOpen() bool {
	return true
}

func (r RenderResultList) Run() {
	r.renderCallback(r.resultList)
}

