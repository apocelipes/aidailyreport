package collector

import (
	"io/fs"
	"os"
	"path/filepath"
)

func FindAllGitRepos(root string, output chan string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		if !isGitRepo(path) {
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

func isGitRepo(path string) bool {
	dir, err := os.Open(filepath.Join(path, ".git"))
	if err != nil {
		return false
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return false
	}

	// a .git dir has at least objects, refs and HEAD
	return len(names) >= 3
}
