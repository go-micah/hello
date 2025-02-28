package main

import "testing"

func TestGetMessage(t *testing.T) {
	got := getMessage()
	want := "Hello, World"

	if got != want {
		t.Errorf("getMessage() = %q, want %q", got, want)
	}
}
