package main

import "testing"

func TestRed(t *testing.T) {
	var r = Red("hello")

	if r != "\033[0;31mhello\033[0m" {
		t.Error("Expected [0;31mhello[0m, got ", r)
	}
}
