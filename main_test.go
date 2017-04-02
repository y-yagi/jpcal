package main

import (
	"bytes"
	"strings"
	"testing"
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
