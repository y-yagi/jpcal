package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	const version = "1.0.0"

	var err error
	var showVersion bool
	var year time.Time
	var calendar Calendar
	var specifyDate string
	var specifyYear string
	var three bool

	date := time.Now()

	flag.StringVar(&specifyDate, "d", "", "Use yyyy-mm as the date.")
	flag.StringVar(&specifyYear, "y", "", "Use yyyy as the year.")
	flag.BoolVar(&three, "3", false, "Display the previous, current and next month surrounding today.")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Println("version:", version)
		os.Exit(0)
		return
	}

	if len(specifyYear) > 0 {
		year, err = time.Parse("2006", specifyYear)
		if err != nil {
			fmt.Printf("Year parse error: %s\n", err)
			os.Exit(1)
		}

		for i := 1; i < 13; i++ {
			date = time.Date(year.Year(), time.Month(i), 1, 0, 0, 0, 0, time.Local)
			calendar.Generate(date)

			if i%3 == 0 {
				calendar.Show()
				calendar.Clear()
			}
		}
	} else if three {
		calendar.Generate(date.AddDate(0, -1, 0))
		calendar.Generate(date)
		calendar.Generate(date.AddDate(0, 1, 0))
		calendar.Show()
	} else {
		if len(specifyDate) > 0 {
			date, err = time.Parse("2006-01", specifyDate)
			if err != nil {
				fmt.Printf("Date parse error: %s\n", err)
				os.Exit(1)
			}
		}

		calendar.Generate(date)
		calendar.Show()
	}
	os.Exit(0)
}
