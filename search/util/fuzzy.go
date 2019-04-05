package util

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"sort"
	"strings"
)

type resultWithScore struct {
	result api.SearchResult
	score float64
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

func GetBestMatches(query string, results []api.SearchResult, minScore float64) []api.SearchResult {

	// Get Scores
	scores := make([]resultWithScore, 0)
	for _, r := range results {

		// Calculate score
		score := getScore(query, r)

		// Check if above minimum
		if score < minScore {
			continue
		}

		// Add result with respective score
		res := resultWithScore{
			result: r,
			score: score,
		}
		scores = append(scores, res)
	}

	// Sort
	sort.Sort(sortableResults(scores))

	// Turn into []api.Result
	results = make([]api.SearchResult, 0)
	for _, s := range scores {
		results = append(results, s.result)
	}

	return results
}

func getScore(query string, result api.SearchResult) float64 {
	query = strings.ToLower(query)
	title := strings.ToLower(result.Title())
	//descr := strings.ToLower(result.Description())

	score := levenshtein.RatioForStrings([]rune(query), []rune(title), levenshtein.DefaultOptions) * 100

	for _, part := range strings.Split(title, " ") {
		if strings.HasPrefix(part, query) {
			score += 40
		}
	}

	return score
}