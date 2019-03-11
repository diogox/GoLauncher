package util

import (
	"github.com/agnivade/levenshtein"
	"github.com/diogox/GoLauncher/api"
	"sort"
)

var queryInput string
type sortableResults []api.Result

func (s sortableResults) Len() int {
	return len(s)
}
func (s sortableResults) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortableResults) Less(i, j int) bool {
	distanceFirst := levenshtein.ComputeDistance(s[i].Title(), queryInput)
	distanceSecond := levenshtein.ComputeDistance(s[j].Title(), queryInput)
	return distanceFirst < distanceSecond
}

func GetBestMatches(query string, results []api.Result) []api.Result {
	queryInput = query
	sort.Sort(sortableResults(results))

	return results
}