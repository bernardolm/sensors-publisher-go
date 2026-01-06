package sqlite

import (
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

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
