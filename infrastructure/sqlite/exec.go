package sqlite

import (
	"context"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

func (c *Client) Exec(ctx context.Context, query string, args ...any) error {
	if c == nil {
		return nil
	}

	_, err := c.db.ExecContext(ctx, query, args)
	if err != nil {
		return errors.Join(err, fmt.Errorf("sqlite: failed to exec query: %s", query))
	}

	return nil
}
