package main

import "time"

func endOfMonth(targetTime time.Time) time.Time {
	return time.Date(targetTime.Year(), targetTime.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

func beginningOfMonth(targetTime time.Time) time.Time {
	return time.Date(targetTime.Year(), targetTime.Month(), 1, 0, 0, 0, 0, time.Local)
}
