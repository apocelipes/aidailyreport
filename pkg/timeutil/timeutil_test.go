package timeutil

import (
	"testing"
	"time"
)

func TestOneDayBefore(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		t        time.Time
		expected time.Time
	}{
		{
			t:        time.Time{},
			expected: time.Time{},
		},
		{
			t:        time.Date(0, 1, 1, 18, 0, 0, 0, tz),
			expected: time.Date(0, 1, 1, 18, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2025, 7, 25, 18, 0, 0, 0, tz),
			expected: time.Date(2025, 7, 24, 21, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2025, 1, 1, 23, 59, 59, 0, tz),
			expected: time.Date(2024, 12, 31, 21, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2024, 3, 1, 12, 0, 0, 0, tz),
			expected: time.Date(2024, 2, 29, 21, 0, 0, 0, tz),
		},
	}
	for _, testCase := range testCases {
		result := OneDayBefore(testCase.t)
		if !result.Equal(testCase.expected) {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}

func TestOneWeekBefore(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		start    time.Weekday
		t        time.Time
		expected time.Time
	}{
		{
			start:    time.Sunday,
			t:        time.Time{},
			expected: time.Time{},
		},
		{
			start:    time.Monday,
			t:        time.Date(0, 1, 5, 18, 0, 0, 0, tz),
			expected: time.Date(0, 1, 5, 18, 0, 0, 0, tz),
		},
		{
			start:    time.Monday,
			t:        time.Date(2025, 7, 25, 18, 0, 0, 0, tz),
			expected: time.Date(2025, 7, 21, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Monday,
			t:        time.Date(2025, 1, 1, 23, 59, 59, 0, tz),
			expected: time.Date(2024, 12, 30, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Monday,
			t:        time.Date(2024, 12, 30, 23, 59, 59, 0, tz),
			expected: time.Date(2024, 12, 30, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Monday,
			t:        time.Date(2024, 3, 1, 12, 0, 0, 0, tz),
			expected: time.Date(2024, 2, 26, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Tuesday,
			t:        time.Date(2024, 12, 30, 23, 59, 59, 0, tz),
			expected: time.Date(2024, 12, 24, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Friday,
			t:        time.Date(2025, 11, 15, 0, 0, 0, 0, tz),
			expected: time.Date(2025, 11, 14, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Sunday,
			t:        time.Date(2025, 11, 14, 23, 59, 59, 0, tz),
			expected: time.Date(2025, 11, 9, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Saturday,
			t:        time.Date(2025, 11, 14, 23, 59, 59, 0, tz),
			expected: time.Date(2025, 11, 8, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Saturday,
			t:        time.Date(2025, 11, 15, 23, 59, 59, 0, tz),
			expected: time.Date(2025, 11, 15, 0, 0, 0, 0, tz),
		},
		{
			start:    time.Saturday,
			t:        time.Date(2025, 11, 16, 23, 59, 59, 0, tz),
			expected: time.Date(2025, 11, 15, 0, 0, 0, 0, tz),
		},
	}
	for _, testCase := range testCases {
		result := OneWeekBefore(testCase.start, testCase.t)
		if !result.Equal(testCase.expected) {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}

func TestCurrentMonthFirstDay(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		t        time.Time
		expected time.Time
	}{
		{
			t:        time.Time{},
			expected: time.Time{},
		},
		{
			t:        time.Date(0, 1, 5, 18, 0, 0, 0, tz),
			expected: time.Date(0, 1, 1, 0, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2025, 7, 31, 18, 0, 0, 0, tz),
			expected: time.Date(2025, 7, 1, 0, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2025, 1, 1, 23, 59, 59, 0, tz),
			expected: time.Date(2025, 1, 1, 0, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2024, 3, 2, 12, 0, 0, 0, tz),
			expected: time.Date(2024, 3, 1, 0, 0, 0, 0, tz),
		},
	}
	for _, testCase := range testCases {
		result := CurrentMonthFirstDay(testCase.t)
		if !result.Equal(testCase.expected) {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}
