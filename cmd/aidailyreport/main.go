package main

import (
	"context"
	"flag"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/apocelipes/aidailyreport/internal/collector"
	"github.com/apocelipes/aidailyreport/internal/data"
	"github.com/apocelipes/aidailyreport/internal/ollama"
	"github.com/apocelipes/aidailyreport/internal/render"
	"github.com/apocelipes/aidailyreport/internal/walker"
	"github.com/apocelipes/aidailyreport/pkg/timeutil"
	"golang.org/x/sync/errgroup"
)

func main() {
	repoPaths := flag.String("path", "", "path to repos")
	isWeekly := flag.Bool("weekly", false, "whether to generate weekly reports")
	authorName := flag.String("name", "", "author name")
	authorEmail := flag.String("email", "", "author email")
	needThinking := flag.Bool("think", false, "turn on/off deep thinking")
	oneLine := flag.Bool("oneline", false, "combine all commits for the same repo into a single line")
	flag.Parse()

	if *repoPaths == "" {
		flag.Usage()
		os.Exit(1)
	}

	since := timeutil.OneDayBefore(time.Now())
	if *isWeekly {
		since = timeutil.OneWeekBefore(time.Now())
	}

	output := make(chan string, runtime.GOMAXPROCS(0)+1)
	commitsPipe := make(chan *data.RepoCommits, runtime.GOMAXPROCS(0)+1)
	wg := &errgroup.Group{}
	for range runtime.GOMAXPROCS(0) {
		wg.Go(func() error {
			for repo := range output {
				commits, err := collector.RecentCommits(repo, &collector.RepoAuthor{
					Name:  *authorName,
					Email: *authorEmail,
				}, since)
				if err != nil {
					return err
				}
				preparedCommits := render.PrepareCommits(commits)
				if preparedCommits != nil {
					commitsPipe <- preparedCommits
				}
			}
			return nil
		})
	}
	go func() {
		err := walker.WalkAllGitRepos(*repoPaths, output)
		if err != nil {
			panic(err)
		}
		close(output)
	}()
	go func() {
		if err := wg.Wait(); err != nil {
			panic(err)
		}
		close(commitsPipe)
	}()

	buff := &strings.Builder{}
	for commits := range commitsPipe {
		err := render.Commits(buff, commits)
		if err != nil {
			panic(err)
		}
	}

	workload := strings.TrimSuffix(buff.String(), "\n")
	err := ollama.SendChatRequest(context.Background(), "qwen3:14b", *needThinking, *oneLine, workload)
	if err != nil {
		panic(err)
	}
}
