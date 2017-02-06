package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func showYearCalendar(specifyYear string) {
	var calendar Calendar
	date := time.Now()

	year, err := time.Parse("2006", specifyYear)
	if err != nil {
		fmt.Printf("Year parse error: %s\n", err)
		os.Exit(1)
	}

	for i := 1; i < 13; i++ {
		date = time.Date(year.Year(), time.Month(i), 1, 0, 0, 0, 0, time.Local)
		calendar.Generate(date)

		if i%3 == 0 {
			calendar.Show(os.Stdout)
			calendar.Clear()
		}
	}
}

func showThreeMonthsCalendar() {
	var calendar Calendar
	date := time.Now()

	calendar.Generate(beginningOfMonth(date).AddDate(0, 0, -1))
	calendar.Generate(date)
	calendar.Generate(endOfMonth(date).AddDate(0, 0, 1))
	calendar.Show(os.Stdout)
}

func showOneMonthCalendar(specifyDate string) {
	var calendar Calendar
	var err error

	date := time.Now()

	if len(specifyDate) > 0 {
		date, err = time.Parse("2006-01", specifyDate)
		if err != nil {
			fmt.Printf("Date parse error: %s\n", err)
			os.Exit(1)
		}
	}

	calendar.Generate(date)
	calendar.Show(os.Stdout)
}

func main() {
	const version = "1.0.0"

	var showVersion bool
	var specifyDate string
	var specifyYear string
	var three bool

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

	if len(flag.Args()) == 1 {
		specifyYear = flag.Args()[0]
	}

	if len(specifyYear) > 0 {
		showYearCalendar(specifyYear)
	} else if three {
		showThreeMonthsCalendar()
	} else {
		showOneMonthCalendar(specifyDate)
	}

	os.Exit(0)
}
