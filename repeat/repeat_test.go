package repeat

import "testing"

func TestRepeat(t *testing.T) {
	r := Repeat("a")
	exp := "aaaaa"

	if r != exp {
		t.Errorf("got %q, want %q", r, exp)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
