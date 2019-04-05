package file

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/result"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const pathRegex = "^((\\/|~)[^/ ]*)+\\/?$"

var _INPUT_ string

func NewFileBrowserMode() *FileBrowserMode {
	regex, err := regexp.Compile(pathRegex)
	if err != nil {
		panic(err)
	}

	return &FileBrowserMode{
		regex: regex,
	}
}

type FileBrowserMode struct {
	launcher *api.Launcher
	regex    *regexp.Regexp
}

func (fb *FileBrowserMode) IsEnabled(input string) bool {
	return fb.regex.MatchString(input)
}

func (fb *FileBrowserMode) HandleInput(input string) api.Action {
	// Set global for helper methods (only used to set query from a relative path, really..)
	_INPUT_ = input

	// Get absolute path
	absPath, _ := getAbsPath(input)
	nOfResults := 9 // TODO: Get number of resutls from preference

	// Not an existing path
	if !isDir(absPath) {

		// Show Suggestions
		return actions.NewRenderResultList(showSuggestions(absPath, nOfResults))
	}

	// If '~', add separator to the end
	if input == "~" {
		return actions.NewSetUserQuery("~/")
	}

	// Otherwise, show results
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		return actions.NewRenderResultList(renderNoMatch())
	}

	return actions.NewRenderResultList(getN(renderResults(files, absPath), nOfResults))
}

func (*FileBrowserMode) DefaultItems(input string) []api.SearchResult {
	return nil
}

func showSuggestions(_path string, nOfResults int) []api.SearchResult {
	fileName := path.Base(_path)
	basePath := _path[:len(_path)-len(fileName)]

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return renderNoMatch()
	}

	allResults := make([]api.SearchResult, 0)
	for _, r := range renderResults(files, basePath) {
		// Check if matches search
		if strings.HasPrefix(strings.ToLower(r.Title()), strings.ToLower(fileName)) {
			allResults = append(allResults, r)
		}
	}

	if len(allResults) == 0 {
		return renderNoMatch()
	}

	return getN(allResults, nOfResults)
}

func renderResults(files []os.FileInfo, basePath string) []api.SearchResult {
	results := make([]api.SearchResult, 0)
	for _, f := range files {

		// Get item's full path
		itemPath := path.Join(basePath, f.Name())

		// Get relative path, if it was typed by the user
		fileName := path.Base(_INPUT_)
		basePath := _INPUT_[:len(_INPUT_)-len(fileName)]

		// Create result
		opts := result.SearchResultOptions{
			Title:            f.Name(),
			Description:      "See what's inside",
			IconPath:         DIR_ICON,
			IsDefaultSelect:  false,
			OnEnterAction:    actions.NewSetUserQuery(path.Join(basePath, f.Name()) + string(os.PathSeparator)),
			OnAltEnterAction: actions.NewOpen(itemPath),
		}

		r := result.NewSearchResult(opts)

		// In case the item is a file
		if !f.IsDir() {
			opts := result.SearchResultOptions{
				Title:            r.Title(),
				Description:      "Open in finder",
				IconPath:         FILE_ICON,
				IsDefaultSelect:  r.IsDefaultSelect(),
				OnEnterAction:    actions.NewOpen(itemPath),
				OnAltEnterAction: actions.NewDoNothing(),
			}

			r = result.NewSearchResult(opts)
		}

		// Add result
		results = append(results, r)
	}

	return results
}

func renderNoMatch() []api.SearchResult {

	opts := result.SearchResultOptions{
		Title:            "No match found!",
		Description:      "Try another path...",
		IconPath:         "warning",
		IsDefaultSelect:  false,
		OnEnterAction:    actions.NewDoNothing(),
		OnAltEnterAction: actions.NewDoNothing(),
	}

	return []api.SearchResult{
		result.NewSearchResult(opts),
	}
}
