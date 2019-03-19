package file

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

const pathRegex = "^((\\/|~)[^/ ]*)+\\/?$"

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
	launcher *api.Launcher
	regex *regexp.Regexp
}

func (fb *FileBrowserMode) IsEnabled(input string) bool {
	if fb.regex.MatchString(input) {
		// Get absolute path, in case it's relative ('~')
		absPath, err := getAbsPath(input)
		if err != nil {
			panic(err)
		}

		_, err = ioutil.ReadDir(absPath)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func (fb *FileBrowserMode) HandleInput(input string) api.Action {
	results := make([]api.Result, 0)

	input, _ = getAbsPath(input)

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

func (*FileBrowserMode) DefaultItems(input string) []api.Result {
	return nil
}

func getAbsPath(pathStr string) (string, error) {
	if !path.IsAbs(pathStr) {
		absPath, err := homedir.Expand(pathStr)
		if err != nil {
			return "", err
		}
		pathStr = absPath
	}
	return pathStr, nil
}