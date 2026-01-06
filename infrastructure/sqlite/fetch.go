package sqlite

import (
	"context"
	"fmt"

	_ "modernc.org/sqlite"
)

func (c *Client) Fetch(ctx context.Context, query string, elements []any) error {
	if c == nil {
		return nil
	}

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("sqlite: select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var row any
		if err := rows.Scan(&row); err != nil {
			return fmt.Errorf("sqlite: scan: %w", err)
		}
		elements = append(elements, row)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("sqlite: rows: %w", err)
	}

	return nil
}
