package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tcnksm/go-holidayjp"
)

var (
	red      = color.New(color.FgRed).SprintFunc()
	blue     = color.New(color.FgBlue, color.Bold).SprintFunc()
	daySpace = "   "
)

func endOfMonth(targetTime time.Time) time.Time {
	return time.Date(targetTime.Year(), targetTime.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

func beginningOfmonth(targetTime time.Time) time.Time {
	return time.Date(targetTime.Year(), targetTime.Month(), 1, 0, 0, 0, 0, time.Local)
}

func printHeader(targetTime time.Time) {
	fmt.Printf("       %d %2d     \n", targetTime.Year(), targetTime.Month())
	fmt.Printf("%s %s %s %s %s %s %s\n", red("日"), "月", "火", "水", "木", "金", blue("土"))
}

func isNeedNewLine(date time.Time) bool {
	return date.Weekday().String() == "Saturday"
}

func decoratedDate(date time.Time) string {
	space := ""
	if date.Day() < 10 {
		space = " "
	}

	_, err := holidayjp.Holiday(date)
	if err == nil {
		return space + red(date.Day())
	} else if date.Weekday().String() == "Saturday" {
		return space + blue(date.Day())
	} else if date.Weekday().String() == "Sunday" {
		return space + red(date.Day())
	} else {
		return space + fmt.Sprint(date.Day())
	}
}

func main() {
	var date time.Time
	now := time.Now()
	printHeader(now)
	firstDate := beginningOfmonth(now)
	lastDate := endOfMonth(now)

	wday := int(firstDate.Weekday())
	fmt.Printf("%s", strings.Repeat(daySpace, wday))

	for i := 1; i < lastDate.Day()+1; i++ {
		date = time.Date(now.Year(), now.Month(), i, 0, 0, 0, 0, time.Local)
		fmt.Printf("%2s ", decoratedDate(date))

		if isNeedNewLine(date) {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}
