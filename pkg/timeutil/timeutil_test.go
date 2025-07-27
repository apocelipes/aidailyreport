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
		t        time.Time
		expected time.Time
	}{
		{
			t:        time.Time{},
			expected: time.Time{},
		},
		{
			t:        time.Date(0, 1, 5, 18, 0, 0, 0, tz),
			expected: time.Date(0, 1, 5, 18, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2025, 7, 25, 18, 0, 0, 0, tz),
			expected: time.Date(2025, 7, 18, 21, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2025, 1, 1, 23, 59, 59, 0, tz),
			expected: time.Date(2024, 12, 25, 21, 0, 0, 0, tz),
		},
		{
			t:        time.Date(2024, 3, 1, 12, 0, 0, 0, tz),
			expected: time.Date(2024, 2, 23, 21, 0, 0, 0, tz),
		},
	}
	for _, testCase := range testCases {
		result := OneWeekBefore(testCase.t)
		if !result.Equal(testCase.expected) {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}
