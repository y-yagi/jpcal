package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tcnksm/go-holidayjp"
)

// Calendar struct
type Calendar struct {
	DateHeader string
	WeekHeader string
	Body       [6]string
}

var (
	red          = color.New(color.FgRed, color.Bold).SprintFunc()
	blue         = color.New(color.FgBlue, color.Bold).SprintFunc()
	white        = color.New(color.FgWhite).SprintFunc()
	reverseVideo = color.New(color.ReverseVideo).SprintFunc()
	daySpace     = "   "
)

func (calendar *Calendar) setHeader(date time.Time) {
	calendar.DateHeader += fmt.Sprintf("     %d年 %02d月       ", date.Year(), date.Month())
	calendar.WeekHeader += fmt.Sprintf("%s %s %s %s %s %s %s   ", red("日"), "月", "火", "水", "木", "金", blue("土"))
}

func (calendar *Calendar) isNeedNewLine(date time.Time) bool {
	return date.Weekday().String() == "Saturday"
}

func (calendar *Calendar) decoratedDate(date time.Time) string {
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

// Generate calendar
func (calendar *Calendar) Generate(date time.Time) {
	var calDate time.Time

	calendar.setHeader(date)
	firstDate := beginningOfMonth(date)
	lastDate := endOfMonth(date)

	wday := int(firstDate.Weekday())
	calendar.Body[0] += fmt.Sprintf("%s", strings.Repeat(daySpace, wday))

	line := 0

	for i := 1; i < lastDate.Day()+1; i++ {
		calDate = time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, time.Local)
		calendar.Body[line] += fmt.Sprintf("%2s ", calendar.decoratedDate(calDate))

		if calendar.isNeedNewLine(calDate) {
			calendar.Body[line] += "  "
			line++
		}
	}

	wday = int(lastDate.Weekday())
	calendar.Body[line] += fmt.Sprintf("%s", strings.Repeat(daySpace, 6-wday))
	calendar.Body[line] += "  "

	for line++; line < len(calendar.Body); line++ {
		calendar.Body[line] += fmt.Sprintf("%s", strings.Repeat(daySpace, 7))
	}
	calendar.Body[line-1] += "  "
}

// Show calendar
func (calendar *Calendar) Show() {
	fmt.Printf(calendar.DateHeader + "\n")
	fmt.Printf(calendar.WeekHeader + "\n")
	for _, element := range calendar.Body {
		fmt.Printf(element + "\n")
	}
}

// Clear calendar
func (calendar *Calendar) Clear() {
	calendar.DateHeader = ""
	calendar.WeekHeader = ""
	for i := range calendar.Body {
		calendar.Body[i] = ""
	}
}
