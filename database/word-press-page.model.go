package database

type wpRendered struct {
	Rendered string `json:"rendered"`
}

type WordPressPage struct {
	Title   wpRendered `json:"title"`
	Content wpRendered `json:"content"`
}
