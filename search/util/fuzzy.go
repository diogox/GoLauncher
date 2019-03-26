package util

import (
	"github.com/agnivade/levenshtein"
	"github.com/diogox/GoLauncher/api"
	"sort"
	"strings"
)

type resultWithScore struct {
	result api.Result
	score int
}

type sortableResults []resultWithScore

func (s sortableResults) Len() int {
	return len(s)
}
func (s sortableResults) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortableResults) Less(i, j int) bool {
	// We want the highest score first
	return s[i].score > s[j].score
}

func GetBestMatches(query string, results []api.Result, minScore int) []api.Result {

	// Get Scores
	scores := make([]resultWithScore, 0)
	for _, r := range results {

		// Calculate score
		score := getScore(query, r)
		
		// Check if above minimum
		if score >= minScore {
			res := resultWithScore{
				result: r,
				score: score,
			}
			scores = append(scores, res)
		}
	}

	// Sort
	sort.Sort(sortableResults(scores))

	// Turn into []api.Result
	results = make([]api.Result, 0)
	for _, s := range scores {
		results = append(results, s.result)
	}

	return results
}

func getScore(query string, result api.Result) int {
	query = strings.ToLower(query)
	title := strings.ToLower(result.Title())
	//descr := strings.ToLower(result.Description())

	score := 0

	score -= levenshtein.ComputeDistance(query, title)

	for _, part := range strings.Split(title, " ") {
		if strings.HasPrefix(part, query) {
			score += 20
		}
	}

	return score
}