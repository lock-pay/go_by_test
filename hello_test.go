package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Axell")
	want := "Hello Axell"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
