package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tcnksm/go-holidayjp"
)

var (
	red          = color.New(color.FgRed, color.Bold).SprintFunc()
	blue         = color.New(color.FgBlue, color.Bold).SprintFunc()
	white        = color.New(color.FgWhite).SprintFunc()
	reverseVideo = color.New(color.ReverseVideo).SprintFunc()
	daySpace     = "   "
)

func endOfMonth(targetTime time.Time) time.Time {
	return time.Date(targetTime.Year(), targetTime.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

func beginningOfMonth(targetTime time.Time) time.Time {
	return time.Date(targetTime.Year(), targetTime.Month(), 1, 0, 0, 0, 0, time.Local)
}

func printHeader(targetTime time.Time) {
	fmt.Printf("       %d %02d     \n", targetTime.Year(), targetTime.Month())
	fmt.Printf("%s %s %s %s %s %s %s\n", red("日"), "月", "火", "水", "木", "金", blue("土"))
}

func isNeedNewLine(date time.Time) bool {
	return date.Weekday().String() == "Saturday"
}

func decoratedDate(date time.Time) string {
	var decoratedDate string
	space := ""
	if date.Day() < 10 {
		space = " "
	}

	_, err := holidayjp.Holiday(date)
	if err == nil {
		decoratedDate = red(date.Day())
	} else if date.Weekday().String() == "Sunday" {
		decoratedDate = red(date.Day())
	} else if date.Weekday().String() == "Saturday" {
		decoratedDate = blue(date.Day())
	} else {
		decoratedDate = white(date.Day())
	}

	today := time.Now()
	if (date.Year() == today.Year()) && (date.Month() == today.Month()) && date.Day() == today.Day() {
		decoratedDate = reverseVideo(decoratedDate)
	}

	return space + decoratedDate
}

func showMonth(date time.Time) {
	var calDate time.Time
	printHeader(date)
	firstDate := beginningOfMonth(date)
	lastDate := endOfMonth(date)

	wday := int(firstDate.Weekday())
	fmt.Printf("%s", strings.Repeat(daySpace, wday))

	for i := 1; i < lastDate.Day()+1; i++ {
		calDate = time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, time.Local)
		fmt.Printf("%2s ", decoratedDate(calDate))

		if isNeedNewLine(calDate) {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func main() {
	var err error
	var year time.Time
	date := time.Now()

	var specifyDate = flag.String("date", "", "Use yyyy-mm as the date.")
	var specifyYear = flag.String("year", "", "Use yyyy as the year.")
	flag.Parse()

	if len(*specifyDate) > 0 {
		date, err = time.Parse("2006-01", *specifyDate)
		if err != nil {
			fmt.Printf("Date parse error: %s\n", err)
			os.Exit(1)
		}
	} else if len(*specifyYear) > 0 {
		year, err = time.Parse("2006", *specifyYear)
		if err != nil {
			fmt.Printf("Year parse error: %s\n", err)
			os.Exit(1)
		}
	}

	if len(*specifyYear) > 0 {
		for i := 1; i < 13; i++ {
			date = time.Date(year.Year(), time.Month(i), 1, 0, 0, 0, 0, time.Local)
			showMonth(date)
			fmt.Printf("\n")
		}
	} else {
		showMonth(date)
	}
	os.Exit(0)
}
