package render

import (
	"io"
	"slices"

	"github.com/apocelipes/aidailyreport/internal/data"
)

func PrepareCommits(commits *data.RepoCommits) *data.RepoCommits {
	if commits == nil || len(commits.Commits) == 0 {
		return nil
	}
	slices.Sort(commits.Commits)
	commits.Commits = slices.Compact(commits.Commits)
	return commits
}

func Commits(dst io.Writer, commits *data.RepoCommits) error {
	if err := render.Execute(dst, commits); err != nil {
		return err
	}
	_, err := dst.Write([]byte{'\n', '\n'})
	return err
}
