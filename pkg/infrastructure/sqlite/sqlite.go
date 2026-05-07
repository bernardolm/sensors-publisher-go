package sqlite

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Client struct {
	db *sql.DB
}

func New(path string) (*Client, error) {
	if path == "" {
		path = "./queue.db"
	}

	if err := ensureDir(path); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		return nil, fmt.Errorf("sqlite: set journal_mode: %w", err)
	}
	if _, err := db.Exec("PRAGMA synchronous=NORMAL;"); err != nil {
		return nil, fmt.Errorf("sqlite: set synchronous: %w", err)
	}
	if _, err := db.Exec("PRAGMA busy_timeout=5000;"); err != nil {
		return nil, fmt.Errorf("sqlite: set busy_timeout: %w", err)
	}

	return &Client{db: db}, nil
}
