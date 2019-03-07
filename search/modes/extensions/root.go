package extensions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/models"
	"github.com/diogox/GoLauncher/search/result"
	"strings"
)

func NewExtensionSearchMode(db *api.DB) *ExtensionSearchMode {

	return & ExtensionSearchMode{
		db: db,
	}
}

type ExtensionSearchMode struct {
	db *api.DB
}

func (esm *ExtensionSearchMode) IsEnabled(input string) bool {
	keyword, _ := getKeywordArgs(input)
	if strings.Contains(input, *keyword + " ") {
		_, err := esm.getExtensionByKeyword(*keyword)
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func (esm *ExtensionSearchMode) HandleInput(input string) api.Action {
	// TODO
	results := make([]api.Result, 0)
	keyword, _ := getKeywordArgs(input)
	if strings.Contains(input, *keyword + " ") {
		extension, _ := esm.getExtensionByKeyword(*keyword)
		action := actions.NewCopyToClipboard(input)
		results = append(results, result.NewSearchResult(extension.Name, extension.Description, extension.IconName, action, action))
	}
	return actions.NewRenderResultList(results)
}

func (*ExtensionSearchMode) DefaultItems(input string) []api.Result {
	return make([]api.Result, 0)
}

func (esm *ExtensionSearchMode) getExtensionByKeyword(keyword string) (models.Extension, error) {
	extension, err := (*esm.db).FindExtensionByKeyword(keyword)
	if err != nil {
		return models.Extension{}, nil
	}
	return extension, nil
}

// TODO: Create a 'query' type for this?
func getKeywordArgs(input string) (*string, []string) {
	components := strings.Split(input, " ")

	if len(components) > 0 {
		return &components[0], components[1:]
	}
	return nil, nil
}
