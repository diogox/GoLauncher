package search

func NewSearch() Search {
	return Search{}
}

type Search struct {}

func (s *Search) HandleInput(input string) []SearchResult {
	results := make([]SearchResult, 0)
	for i := 0; i < 4; i++ {
		results = append(results, SearchResult{
			title: input,
			description: "Description: " + input,
			iconPath: "/",
		})
	}

	return results
}

type SearchResult struct {
	iconPath string
	title string
	description string
}

func (r SearchResult) Title() string {
	return r.title
}

func (r SearchResult) Description() string {
	return r.description
}

func (r SearchResult) IconPath() string {
	return r.iconPath
}

