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

// snippetModel which wraps the sql db connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// returns sql result type which includes basic info about what happened
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// get the last item inserted's id and return it
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// id has a type of int64 needs to be converted to int
	return int(id), nil
}

// return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// this will reeturn the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
