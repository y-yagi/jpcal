package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func showYearCalendar(specifyYear string, w io.Writer) {
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
			calendar.Show(w)
			calendar.Clear()
		}
	}
}

func showThreeMonthsCalendar(w io.Writer) {
	var calendar Calendar
	date := time.Now()

	calendar.Generate(beginningOfMonth(date).AddDate(0, 0, -1))
	calendar.Generate(date)
	calendar.Generate(endOfMonth(date).AddDate(0, 0, 1))
	calendar.Show(w)
}

func showOneMonthCalendar(specifyDate string, w io.Writer) {
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
	calendar.Show(w)
}

func run(args []string, out, err io.Writer) int {
	const version = "1.0.0"

	var showVersion bool
	var specifyDate string
	var specifyYear string
	var three bool

	flags := flag.NewFlagSet("jpcal", flag.ExitOnError)
	flags.SetOutput(err)
	flags.StringVar(&specifyDate, "d", "", "Use yyyy-mm as the date.")
	flags.StringVar(&specifyYear, "y", "", "Use yyyy as the year.")
	flags.BoolVar(&three, "3", false, "Display the previous, current and next month surrounding today.")
	flags.BoolVar(&showVersion, "v", false, "show version")
	flags.Parse(args[1:])

	if showVersion {
		fmt.Fprintln(out, "version:", version)
		return 0
	}

	if len(flag.Args()) == 1 {
		specifyYear = flag.Args()[0]
	}

	if len(specifyYear) > 0 {
		showYearCalendar(specifyYear, out)
	} else if three {
		showThreeMonthsCalendar(out)
	} else {
		showOneMonthCalendar(specifyDate, out)
	}

	return 0
}

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}
