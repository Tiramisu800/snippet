package models

import (
	"context"
	"errors"
	"fmt"
	//"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet.(like in postgres)
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	q := "INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, NOW(), (NOW() + INTERVAL '1 DAY' * $3)) RETURNING id"

	row := m.DB.QueryRow(context.Background(),
		q,
		title, content, expires)

	var id uint64

	err := row.Scan(&id)
	if err != nil {
		fmt.Printf("Unable to INSERT: %v\n", err)
		return 0, err
	}

	fmt.Println(id)

	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	row := m.DB.QueryRow(context.Background(),
		"SELECT id, title, content, created, expires FROM snippets WHERE expires > NOW() AND id = $1", id)
	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// This will return the 10 most recently created snippets.

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
			WHERE expires > NOW() ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
