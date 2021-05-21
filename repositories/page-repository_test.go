package repositories

import (
	"database/sql"
	"log"
	"testing"

	r "bitbucket.org/y16i/backend-go/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var p = &r.Page{
	Title:   r.Rendered{Rendered: "title"},
	Content: r.Rendered{Rendered: "content"},
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error %s wasn't expected", err)
	}
	return db, mock
}

func TestFindBySlug(t *testing.T) {
	db, mock := NewMock()
	repo := NewPageRepository(db)
	defer func() {
		repo.Close()
	}()

	query := "select post_title, post_content from wp_posts where post_type = 'page' and post_name=?"
	rows := sqlmock.NewRows([]string{"post_title", "post_content"}).
		AddRow(p.Title.Rendered, p.Content.Rendered)
	mock.ExpectQuery(query).WithArgs("some-slug").WillReturnRows(rows)

	page, err := repo.FindBySlug("some-slug")
	assert.NotNil(t, page)
	assert.NoError(t, err)
}
