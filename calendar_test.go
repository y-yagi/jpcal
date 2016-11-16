package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	var calendar Calendar
	date, _ := time.Parse("2006-01-02", "2016-09-01")
	calendar.Generate(date)

	expect := "     2016年 09月       "
	if calendar.DateHeader != expect {
		t.Errorf("Expect calendar.DateHeader is %s, but %s", expect, calendar.DateHeader)
	}

	expect = fmt.Sprintf("%s %s %s %s %s %s %s   ", red("日"), "月", "火", "水", "木", "金", blue("土"))
	if calendar.WeekHeader != expect {
		t.Errorf("Expect calendar.WeekHeader is %s, but %s", expect, calendar.WeekHeader)
	}

	expect = fmt.Sprintf("%s %s %s %s  %s  %s  %s   ", "  ", "  ", "  ", "  ", "1", "2", blue("3"))
	if calendar.Body[0] != expect {
		t.Errorf("Expect calendar.Body[0] is %s, but %s", expect, calendar.Body[0])
	}

	expect = fmt.Sprintf(" %s  %s  %s  %s  %s  %s %s   ", red("4"), "5", "6", "7", "8", "9", blue("10"))
	if calendar.Body[1] != expect {
		t.Errorf("Expect calendar.Body[1] is %s, but %s", expect, calendar.Body[1])
	}

	expect = fmt.Sprintf("%s %s %s %s %s %s %s   ", red("11"), "12", "13", "14", "15", "16", blue("17"))
	if calendar.Body[2] != expect {
		t.Errorf("Expect calendar.Body[2] is %s, but %s", expect, calendar.Body[2])
	}

	// 2016-09-19 and 09-22 is holiday
	expect = fmt.Sprintf("%s %s %s %s %s %s %s   ", red("18"), red("19"), "20", "21", red("22"), "23", blue("24"))
	if calendar.Body[3] != expect {
		t.Errorf("Expect calendar.Body[3] is %s, but %s", expect, calendar.Body[3])
	}

	expect = fmt.Sprintf("%s %s %s %s %s %s %s   ", red("25"), "26", "27", "28", "29", "30", "  ")
	if calendar.Body[4] != expect {
		t.Errorf("Expect calendar.Body[4] is %s, but %s", expect, calendar.Body[4])
	}
}

func TestClear(t *testing.T) {
	var calendar Calendar

	date, _ := time.Parse("2006-01-02", "2016-10-08")
	calendar.Generate(date)
	calendar.Clear()

	if len(calendar.DateHeader) > 0 {
		t.Errorf("Expect calendar.DateHeader is empty, but %s", calendar.DateHeader)
	}

	if len(calendar.WeekHeader) > 0 {
		t.Errorf("Expect calendar.WeekHeader is empty, but %s", calendar.WeekHeader)
	}

	for _, element := range calendar.Body {
		if len(element) > 0 {
			t.Errorf("Expect calendar.Body is empty, but %s", element)
		}
	}
}
