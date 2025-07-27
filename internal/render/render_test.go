package render

import (
	"reflect"
	"testing"

	"github.com/apocelipes/aidailyreport/internal/data"
)

func TestPrepareCommits(t *testing.T) {
	testCases := []struct {
		commits         *data.RepoCommits
		expectedCommits *data.RepoCommits
	}{
		{
			commits:         nil,
			expectedCommits: nil,
		},
		{
			commits:         &data.RepoCommits{RepoName: "test"},
			expectedCommits: nil,
		},
		{
			commits:         &data.RepoCommits{},
			expectedCommits: nil,
		},
		{
			commits: &data.RepoCommits{
				RepoName: "",
				Commits: []string{
					"fix: abc",
					"feat: def",
					"perf: ghi",
				},
			},
			expectedCommits: &data.RepoCommits{
				RepoName: "",
				Commits: []string{
					"feat: def",
					"fix: abc",
					"perf: ghi",
				},
			},
		},
		{
			commits: &data.RepoCommits{
				RepoName: "",
				Commits: []string{
					"fix: abc",
					"feat: def",
					"feat: def",
					"perf: ghi",
					"fix: abc",
				},
			},
			expectedCommits: &data.RepoCommits{
				RepoName: "",
				Commits: []string{
					"feat: def",
					"fix: abc",
					"perf: ghi",
				},
			},
		},
	}
	for _, testCase := range testCases {
		result := PrepareCommits(testCase.commits)
		if !reflect.DeepEqual(result, testCase.expectedCommits) {
			t.Errorf("PrepareCommits error, want %#v, got %#v", testCase.expectedCommits, result)
		}
	}
}
