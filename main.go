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

// CalendarLine is calendar line number
const CalendarLine = 8

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

func setHeader(targetTime time.Time, calendar *[CalendarLine]string) {
	calendar[0] += fmt.Sprintf("     %d年 %02d月       ", targetTime.Year(), targetTime.Month())
	calendar[1] += fmt.Sprintf("%s %s %s %s %s %s %s   ", red("日"), "月", "火", "水", "木", "金", blue("土"))
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

func setCalendar(date time.Time, calendar *[CalendarLine]string) {
	var calDate time.Time

	line := 2

	setHeader(date, calendar)
	firstDate := beginningOfMonth(date)
	lastDate := endOfMonth(date)

	wday := int(firstDate.Weekday())
	calendar[line] += fmt.Sprintf("%s", strings.Repeat(daySpace, wday))

	for i := 1; i < lastDate.Day()+1; i++ {
		calDate = time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, time.Local)
		calendar[line] += fmt.Sprintf("%2s ", decoratedDate(calDate))

		if isNeedNewLine(calDate) {
			calendar[line] += "  "
			line++
		}
	}

	wday = int(lastDate.Weekday())
	calendar[line] += fmt.Sprintf("%s", strings.Repeat(daySpace, 6-wday))
	calendar[line] += "  "

	for line++; line < CalendarLine; line++ {
		calendar[line] += fmt.Sprintf("%s", strings.Repeat(daySpace, 7))
	}
	calendar[line-1] += "  "
}

func showCalendar(calendar *[CalendarLine]string) {
	for i, element := range calendar {
		fmt.Printf(element + "\n")
		calendar[i] = ""
	}
}

func main() {
	var err error
	var year time.Time
	date := time.Now()
	var calendar [CalendarLine]string

	var specifyDate = flag.String("d", "", "Use yyyy-mm as the date.")
	var specifyYear = flag.String("y", "", "Use yyyy as the year.")
	var three = flag.Bool("3", false, "Display the previous, current and next month surrounding today.")
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
			setCalendar(date, &calendar)

			if i%3 == 0 {
				showCalendar(&calendar)
			}
		}
	} else if *three {
		setCalendar(date.AddDate(0, -1, 0), &calendar)
		setCalendar(date, &calendar)
		setCalendar(date.AddDate(0, 1, 0), &calendar)
		showCalendar(&calendar)
	} else {
		setCalendar(date, &calendar)
		showCalendar(&calendar)
	}
	os.Exit(0)
}
