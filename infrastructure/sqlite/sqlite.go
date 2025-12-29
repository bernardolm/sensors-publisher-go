package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

const defaultMaxBatch = 10

type Queue struct {
	db       *sql.DB
	maxBatch int
}

func New(path string, maxBatch int) (*Queue, error) {
	if path == "" {
		return nil, errors.New("sqlite: path is required")
	}

	if maxBatch <= 0 {
		maxBatch = defaultMaxBatch
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

	schema := `CREATE TABLE IF NOT EXISTS queue (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		target TEXT NOT NULL,
		topic TEXT NOT NULL DEFAULT '',
		payload BLOB NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("sqlite: create table: %w", err)
	}

	if _, err := db.Exec("CREATE INDEX IF NOT EXISTS queue_target_id ON queue(target, id);"); err != nil {
		return nil, fmt.Errorf("sqlite: create index: %w", err)
	}

	return &Queue{db: db, maxBatch: maxBatch}, nil
}

func ensureDir(path string) error {
	if path == ":memory:" || strings.HasPrefix(path, "file:") {
		return nil
	}

	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}

	return os.MkdirAll(dir, 0o755)
}

func (q *Queue) Enqueue(ctx context.Context, target, topic string, payload []byte) error {
	if q == nil {
		return nil
	}

	if target == "" {
		return errors.New("sqlite: target is required")
	}

	_, err := q.db.ExecContext(ctx,
		"INSERT INTO queue (target, topic, payload) VALUES (?, ?, ?)",
		target,
		topic,
		payload,
	)
	if err != nil {
		return fmt.Errorf("sqlite: insert: %w", err)
	}

	return nil
}

func (q *Queue) Flush(ctx context.Context, target string, send func(context.Context, string, []byte) error) (int, error) {
	if q == nil {
		return 0, nil
	}

	if target == "" {
		return 0, errors.New("sqlite: target is required")
	}

	if send == nil {
		return 0, errors.New("sqlite: send func is required")
	}

	rows, err := q.db.QueryContext(ctx,
		"SELECT id, topic, payload FROM queue WHERE target = ? ORDER BY id LIMIT ?",
		target,
		q.maxBatch,
	)
	if err != nil {
		return 0, fmt.Errorf("sqlite: select: %w", err)
	}
	defer rows.Close()

	type entry struct {
		id      int64
		topic   string
		payload []byte
	}

	entries := []entry{}

	for rows.Next() {
		var e entry
		if err := rows.Scan(&e.id, &e.topic, &e.payload); err != nil {
			return 0, fmt.Errorf("sqlite: scan: %w", err)
		}
		entries = append(entries, e)
	}

	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("sqlite: rows: %w", err)
	}

	sent := 0
	for _, e := range entries {
		if err := send(ctx, e.topic, e.payload); err != nil {
			return sent, err
		}
		if _, err := q.db.ExecContext(ctx, "DELETE FROM queue WHERE id = ?", e.id); err != nil {
			return sent, fmt.Errorf("sqlite: delete: %w", err)
		}
		sent++
	}

	return sent, nil
}

func (q *Queue) Close() error {
	if q == nil || q.db == nil {
		return nil
	}

	return q.db.Close()
}
