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
	results := make([]api.SearchResult, 0)

	calcResult, err := c.calculate(input)
	if err != nil {

		opts := result.SearchResultOptions{
			Title:            "Error!",
			Description:      err.Error(),
			IconPath:         "calc",
			IsDefaultSelect:  true,
			OnEnterAction:    actions.NewDoNothing(),
			OnAltEnterAction: actions.NewDoNothing(),
		}

		r := result.NewSearchResult(opts)
		results = append(results, r)
		return actions.NewRenderResultList(results)
	}

	opts := result.SearchResultOptions{
		Title:            calcResult,
		Description:      "Copy to Clipboard",
		IconPath:         "calc",
		IsDefaultSelect:  true,
		OnEnterAction:    actions.NewCopyToClipboard(calcResult),
		OnAltEnterAction: actions.NewCopyToClipboard(calcResult),
	}

	r := result.NewSearchResult(opts)
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

func (*CalcSearchMode) DefaultItems(input string) []api.SearchResult {
	return nil
}
