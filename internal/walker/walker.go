package walker

import (
	"io/fs"
	"os"
	"path/filepath"
)

func WalkAllGitRepos(root string, output chan string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		if _, err := os.Stat(filepath.Join(path, ".git")); err != nil {
			return nil
		}
		abs, absErr := filepath.Abs(path)
		if absErr != nil {
			return absErr
		}
		output <- abs
		return nil
	})
}
