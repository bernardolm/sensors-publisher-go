package sqlite

// import (
// 	"context"

// 	infrastructuresqlite "github.com/bernardolm/sensors-publisher-go/infrastructure/sqlite"
// )

// const (
// 	queryCreate = `
// 		CREATE TABLE IF NOT EXISTS messages (
// 			id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			target TEXT NOT NULL,
// 			topic TEXT NOT NULL DEFAULT '',
// 			payload BLOB NOT NULL,
// 			created_at DATETIME NOT NULL,
// 			sent_at DATETIME
// 		);

// 		CREATE INDEX IF NOT EXISTS messages_target_sent_id ON messages(target, sent_at, id);
// 	`
// )

// func (c *sqlite) Create(ctx context.Context) error {
// 	err :=  c.   infrastructuresqlite.Exec(ctx, queryCreate)
// 	return err
// }
