package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// Client owns a reusable PostgreSQL connection pool.
type Client struct {
	db *sql.DB
}

// New opens a PostgreSQL connection pool without requiring initial connectivity.
func New(_ context.Context, dsn string) (*Client, error) {
	configuration, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: parse connection: %w", err)
	}
	configuration.RuntimeParams["timezone"] = "America/Sao_Paulo"
	db := stdlib.OpenDB(*configuration)

	db.SetConnMaxIdleTime(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxOpenConns(100)

	return &Client{db: db}, nil
}

// DB returns the underlying database connection pool.
func (c *Client) DB() *sql.DB {
	return c.db
}

// Close closes the PostgreSQL connection pool.
func (c *Client) Close() error {
	if c == nil || c.db == nil {
		return nil
	}

	return c.db.Close()
}
