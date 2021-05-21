package models

type Rendered struct {
	Rendered string `json:"rendered"`
}

type Page struct {
	Title   Rendered `json:"title"`
	Content Rendered `json:"content"`
}

type PageRepository interface {
	FindBySlug(slug string) (*Page, error)
}
