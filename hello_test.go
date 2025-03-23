package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("it says the name", func(t *testing.T) {
		got := Hello("Axell")
		want := "Hello, Axell"

		AssertCorrectness(t, got, want)

	})

	t.Run("it says world if no name is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		AssertCorrectness(t, got, want)
	})
}

func AssertCorrectness(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
