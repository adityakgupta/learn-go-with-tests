package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("chris", "en")
		want := "Hello, chris"

		assertMessage(t, got, want)
	})

	t.Run("say 'Hello, World' by default", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
