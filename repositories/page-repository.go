package repositories

import (
	"database/sql"
	"log"

	"bitbucket.org/y16i/backend-go/models"
)

// Repository for Page model
type PageRepository struct {
	db *sql.DB
}

func NewPageRepository(db *sql.DB) *PageRepository {
	return &PageRepository{
		db: db,
	}
}

func (r PageRepository) Close() {
	r.db.Close()
}

func (r *PageRepository) FindBySlug(slug string) (*models.Page, error) {
	// query: select post_title, post_content from wp_posts where post_type = 'page' and post_name = 'portfolio-summary';
	if r == nil {
		log.Fatal("DB connection is nil")
	}
	row := r.db.QueryRow("select post_title, post_content from wp_posts where post_type = 'page' and post_name=?", slug)
	page := new(models.Page)
	queryErr := row.Scan(&page.Title.Rendered, &page.Content.Rendered)
	if queryErr != nil {
		log.Println(queryErr)
	}

	return page, queryErr
}
