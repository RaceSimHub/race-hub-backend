package model

type ListTemplateData struct {
	Title     string
	NewURL    string
	SearchURL string
	EditURL   string
	DeleteURL string
	Columns   []string
	Items     []map[string]any
	Search    string
	Offset    int
	Limit     int
	Total     int
	ExtraPage int
}

func (l ListTemplateData) ColumnsCount() int {
	return len(l.Columns) + 1
}
