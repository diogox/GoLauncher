package shortcut

import (
	"fmt"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/api/models"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/diogox/GoLauncher/sqlite"
	"regexp"
	"strings"
)

func NewShortcutSearchMode(db *api.DB) ShortcutSearchMode {

	sqlite.LoadDefaultShortcuts(db)

	return ShortcutSearchMode{
		db: db,
	}
}

type ShortcutSearchMode struct {
	db *api.DB
}

func (ssm ShortcutSearchMode) IsEnabled(input string) bool {

	if shortcut := ssm.getActiveShortcut(input); shortcut != nil {
		return true
	}
	return false
}

func (ssm ShortcutSearchMode) HandleInput(input string) api.Action {

	shortcut := ssm.getActiveShortcut(input)
	if shortcut == nil {
		panic("No active shortcut!")
	}

	pattern := fmt.Sprintf("^%s ", shortcut.Keyword)
	isMatch, err := regexp.Match(pattern, []byte(input))
	if err != nil {
		panic(err)
	}

	if isMatch {
		input = strings.Replace(input, shortcut.Keyword + " ", "", 1)
	}

	results := make([]api.Result, 0)

	descr := strings.Replace(shortcut.Cmd, "%s", input, -1)
	action := actions.NewOpenUrl(descr)
	r := result.NewSearchResult(shortcut.Name, descr, shortcut.IconName, action, action)

	results = append(results, r)
	return actions.NewRenderResultList(results)
}

func (ssm ShortcutSearchMode) getActiveShortcut(input string) *models.ShortcutInfo {
	shortcuts, err := (*ssm.db).GetAllShortcuts()
	if err != nil {
		panic(err)
	}

	for _, shortcut := range shortcuts {
		if shortcut.IsActive && strings.HasPrefix(input, shortcut.Keyword + " ") {
			return &shortcut
		}
	}

	return nil
}

