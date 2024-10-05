package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chetan")

	got := buffer.String()
	want := "Hello, Chetan"

	if got != want {
		t.Errorf("Expected %s, but got %s", want, got)
	}
}
