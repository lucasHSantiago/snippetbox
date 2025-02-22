package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// Inset a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Get a specific snippet by id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Return the last 10 most recent created snippets
func (m *SnippetModel) Latest() ([]*SnippetModel, error) {
	return nil, nil
}
