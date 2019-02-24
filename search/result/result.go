package result

func NewSearchResult(title string, descr string, iconPath string) SearchResult {
	return SearchResult{
		title: title,
		description: descr,
		iconPath: iconPath,
	}
}

type SearchResult struct {
	title string
	description string
	iconPath string
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

