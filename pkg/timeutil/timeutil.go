package timeutil

import "time"

func OneDayBefore(t time.Time) time.Time {
	res := t.AddDate(0, 0, -1)
	res = time.Date(res.Year(), res.Month(), res.Day(), 21, 0, 0, 0, t.Location())
	if res.UnixNano() <= 0 {
		return t
	}
	return res
}

func OneWeekBefore(t time.Time) time.Time {
	res := t.AddDate(0, 0, -7)
	res = time.Date(res.Year(), res.Month(), res.Day(), 21, 0, 0, 0, t.Location())
	if res.UnixNano() <= 0 {
		return t
	}
	return res
}

func CurrentMonthFirstDay(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	res := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return res
}
