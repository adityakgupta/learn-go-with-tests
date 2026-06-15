package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "World")

	got := buf.String()
	want := "Hello World"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
