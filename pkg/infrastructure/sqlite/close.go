package sqlite

import (
	_ "modernc.org/sqlite"
)

func (c *Client) Close() error {
	if c == nil || c.db == nil {
		return nil
	}

	return c.db.Close()
}
