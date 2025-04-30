package dto

type SearchOptions struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

type SortOption struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type FilterOption struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}
