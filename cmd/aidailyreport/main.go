package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/apocelipes/aidailyreport/internal/collector"
	"github.com/apocelipes/aidailyreport/internal/data"
	"github.com/apocelipes/aidailyreport/internal/ollama"
	"github.com/apocelipes/aidailyreport/internal/render"
	"github.com/apocelipes/aidailyreport/pkg/timeutil"
	"golang.org/x/sync/errgroup"
)

func main() {
	repoPaths := flag.String("path", "", "path to repos")
	isWeekly := flag.Bool("weekly", false, "whether to generate weekly reports")
	isMonthly := flag.Bool("monthly", false, "whether to generate monthly reports")
	authorName := flag.String("name", "", "author name")
	authorEmail := flag.String("email", "", "author email")
	needThinking := flag.Bool("think", false, "turn on/off deep thinking")
	oneLine := flag.Bool("oneline", false, "combine all commits for the same repo into a single line")
	modelName := flag.String("model", "qwen3:14b", "the LLM will be used")
	flag.Parse()

	if *repoPaths == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *isWeekly && *isMonthly {
		_, _ = fmt.Fprintln(os.Stderr, "cannot both set -weekly and -monthly")
		flag.Usage()
		os.Exit(1)
	}

	now := time.Now()
	since := timeutil.OneDayBefore(now)
	if *isWeekly {
		since = timeutil.OneWeekBefore(now)
	} else if *isMonthly {
		since = timeutil.CurrentMonthFirstDay(now)
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
		err := collector.FindAllGitRepos(*repoPaths, output)
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

	if buff.Len() == 0 {
		fmt.Fprintln(os.Stderr, "No content!")
		os.Exit(1)
	}

	workload := strings.TrimSuffix(buff.String(), "\n")
	err := ollama.SendChatRequest(context.Background(), *modelName, *needThinking, *oneLine, workload)
	if err != nil {
		panic(err)
	}
}
