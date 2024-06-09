package search

type TextSearch struct {
	dto *TextSearchFlagDTO
}

func NewTextSearch(dto *TextSearchFlagDTO) *TextSearch {
	return &TextSearch{dto}
}

func (receiver TextSearch) Search(rows []string) []string {
	return nil
}
