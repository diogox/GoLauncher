package calc

import (
	"github.com/diogox/GoLauncher/common"
	"github.com/diogox/GoLauncher/common/actions"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/relnod/calcgo"
	"regexp"
	"strconv"
	"strings"
)

const calcRegexString = "([-+]?[0-9]*\\.?[0-9]+[\\/\\+\\-\\*])+([-+]?[0-9]*\\.?[0-9]+)"

func NewCalcSearchMode(launcher *common.Launcher) *CalcSearchMode {
	regex, err := regexp.Compile(calcRegexString)
	if err != nil {
		panic(err)
	}

	return &CalcSearchMode{
		launcher: launcher,
		calcRegex: regex,
	}
}

type CalcSearchMode struct {
	launcher *common.Launcher
	calcRegex *regexp.Regexp
}

func (c *CalcSearchMode) IsEnabled(input string) bool {
	return c.calcRegex.MatchString(input)
}

func (c *CalcSearchMode) HandleInput(input string) common.Action {
	results := make([]common.Result, 0)

	calcResult := c.calculate(input)

	action := actions.NewCopyToClipboardAction(calcResult)
	r := result.NewSearchResult(calcResult, "Copy to Clipboard", "calc", action, action)
	results = append(results, r)

	return actions.NewRenderResultListAction(results, (*c.launcher).ShowResults)
}

func (*CalcSearchMode) calculate(input string) string {
	// TODO: Refactor this code
	input = strings.Replace(input, "+", " + ", -1)
	input = strings.Replace(input, "-", " - ", -1)
	input = strings.Replace(input, "*", " * ", -1)
	input = strings.Replace(input, "/", " / ", -1)

	result, _ := calcgo.Calc(input)
	resultStr := strconv.FormatFloat(result, 'f', -1, 32)
	return resultStr
}

