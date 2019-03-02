package file

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/common/actions"
	"github.com/diogox/GoLauncher/search/result"
	"io/ioutil"
	"regexp"
	"strings"
)

const pathRegex = "^(/[^/ ]*)+/?$"

func NewFileBrowserMode() *FileBrowserMode {
	regex, err := regexp.Compile(pathRegex)
	if err != nil {
		panic(err)
	}

	return &FileBrowserMode {
		regex: regex,
	}
}

type FileBrowserMode struct {
	launcher *common.Launcher
	regex *regexp.Regexp
}

func (fb *FileBrowserMode) IsEnabled(input string) bool {
	if fb.regex.MatchString(input) {
		_, err := ioutil.ReadDir(input)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func (fb *FileBrowserMode) HandleInput(input string) common.Action {
	results := make([]common.Result, 0)

	files, err := ioutil.ReadDir(input)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		if i > 4 {
			break
		}

		var path string
		if strings.HasSuffix(input, "/") {
			path = input + file.Name()
		} else {
			path = input + "/" + file.Name()
		}

		action := actions.NewOpen(path)
		r := result.NewSearchResult(file.Name(), "Open in finder", "nemo", action, action)
		results = append(results, r)
	}

	return actions.NewRenderResultList(results)
}