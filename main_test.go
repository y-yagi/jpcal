package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRunVersion(t *testing.T) {
	out, err := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("jpcal -v", " ")

	status := run(args, out, err)
	if status != 0 {
		t.Errorf("Expect status is 0, but %d", status)
	}

	expected := "version: 1.0.0"
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}
}

func TestRunWithoutArgs(t *testing.T) {
	out, err := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("jpcal", " ")

	status := run(args, out, err)
	if status != 0 {
		t.Errorf("Expect status is 0, but %d", status)
	}

	date := time.Now()
	expected := fmt.Sprintf("     %d年 %02d月       ", date.Year(), date.Month())
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}
}

func TestRunWithSpecifyYear(t *testing.T) {
	out, err := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("jpcal -y 2017", " ")

	status := run(args, out, err)
	if status != 0 {
		t.Errorf("Expect status is 0, but %d", status)
	}

	expected := "2017年 01月            2017年 02月            2017年 03月"
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}
}

func TestYearOptionsShowCurrentYearByDefault(t *testing.T) {
	out, err := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("jpcal -y", " ")

	status := run(args, out, err)
	if status != 0 {
		t.Errorf("Expect status is 0, but %d", status)
	}

	date := time.Now()
	expected := fmt.Sprintf("%d年 %02d月", date.Year(), date.Month())
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}
}

func TestRunWithThree(t *testing.T) {
	out, err := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("jpcal -3", " ")

	status := run(args, out, err)
	if status != 0 {
		t.Errorf("Expect status is 0, but %d", status)
	}

	date := time.Now()
	expected := fmt.Sprintf("%d年 %02d月", date.Year(), date.Month())
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}

	lastMonth := beginningOfMonth(date).AddDate(0, 0, -1)
	expected = fmt.Sprintf("%d年 %02d月", lastMonth.Year(), lastMonth.Month())
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}

	nextMonth := beginningOfMonth(date).AddDate(0, 0, 1)
	expected = fmt.Sprintf("%d年 %02d月", nextMonth.Year(), nextMonth.Month())
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}
}

func TestRunWithSpecifyDate(t *testing.T) {
	out, err := new(bytes.Buffer), new(bytes.Buffer)
	args := strings.Split("jpcal -d 2017-04", " ")

	status := run(args, out, err)
	if status != 0 {
		t.Errorf("Expect status is 0, but %d", status)
	}

	expected := "2017年 04月"
	if !strings.Contains(out.String(), expected) {
		t.Errorf("Expect out is %q, but %q", out.String(), expected)
	}
}
