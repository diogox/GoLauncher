package calc

import (
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
	"github.com/diogox/GoLauncher/search/result"
	"github.com/relnod/calcgo"
	"regexp"
	"strconv"
	"strings"
)

const calcRegexString = "([-+]?[0-9]*\\.?[0-9]+[\\/\\+\\-\\*])+([-+]?[0-9]*\\.?[0-9]+)"

func NewCalcSearchMode() *CalcSearchMode {
	regex, err := regexp.Compile(calcRegexString)
	if err != nil {
		panic(err)
	}

	return &CalcSearchMode{
		calcRegex: regex,
	}
}

type CalcSearchMode struct {
	calcRegex *regexp.Regexp
}

func (c *CalcSearchMode) IsEnabled(input string) bool {
	return c.calcRegex.MatchString(input)
}

func (c *CalcSearchMode) HandleInput(input string) api.Action {
	results := make([]api.Result, 0)

	calcResult, err := c.calculate(input)
	if err != nil {
		action := actions.NewDoNothing()
		r := result.NewSearchResult("Error!", err.Error(), "calc", true, action, action)
		results = append(results, r)
		return actions.NewRenderResultList(results)
	}

	action := actions.NewCopyToClipboard(calcResult)
	r := result.NewSearchResult(calcResult, "Copy to Clipboard", "calc", true, action, action)
	results = append(results, r)

	return actions.NewRenderResultList(results)
}

func (*CalcSearchMode) calculate(input string) (string, error) {
	// TODO: Refactor this code
	input = strings.Replace(input, "+", " + ", -1)
	input = strings.Replace(input, "-", " - ", -1)
	input = strings.Replace(input, "*", " * ", -1)
	input = strings.Replace(input, "/", " / ", -1)

	mathRes, errs := calcgo.Calc(input)
	if len(errs) != 0 {
		return "", errors.New("Invalid Expression!")
	}

	resultStr := strconv.FormatFloat(mathRes, 'f', -1, 32)
	return resultStr, nil
}

func (*CalcSearchMode) DefaultItems(input string) []api.Result {
	return nil
}