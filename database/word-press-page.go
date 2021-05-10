package database

import "log"

func GetWordPressPage(slug string) (*WordPressPage, error) {
	// query: select post_title, post_content from wp_posts where post_type = 'page' and post_name = 'portfolio-summary';
	if conn == nil {
		log.Fatal("conn is nil")
	}
	row := conn.QueryRow("select post_title, post_content from wp_posts where post_type = 'page' and post_name=?", slug)
	page := new(WordPressPage)
	queryErr := row.Scan(&page.Title.Rendered, &page.Content.Rendered)
	if queryErr != nil {
		log.Println(queryErr)
	}

	return page, queryErr
}
