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
