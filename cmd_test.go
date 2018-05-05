package main

import (
	"testing"
)

// TestCmd just assumes that git can be executed
func TestCmd(t *testing.T) {
	_, err := Cmd("./", "--help")

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
}
