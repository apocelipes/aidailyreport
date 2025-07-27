package collector

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/apocelipes/aidailyreport/internal/data"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type RepoAuthor struct {
	Name  string
	Email string
}

func RecentCommits(repoPath string, author *RepoAuthor, since time.Time) (*data.RepoCommits, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	headRef, err := repo.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := repo.Log(&git.LogOptions{
		From:  headRef.Hash(),
		Since: &since,
	})
	if err != nil {
		return nil, err
	}

	var messages []string
	err = cIter.ForEach(func(c *object.Commit) error {
		if c.Author.When.Before(since) {
			return nil
		}
		if author.Email != "" && c.Author.Email != author.Email {
			return nil
		}
		if author.Name != "" && c.Author.Name != author.Name {
			return nil
		}
		messages = append(messages, strings.TrimSuffix(c.Message, "\n"))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data.RepoCommits{
		RepoName: filepath.Base(repoPath),
		Commits:  messages,
	}, nil
}
