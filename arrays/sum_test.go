package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("n numbers", func(t *testing.T) {
		a := []int{1, 2, 3, 4}

		got := Sum(a)
		want := 10

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	check := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("check tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		check(t, got, want)
	})
	t.Run("check empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{})
		want := []int{2, 9, 0}

		check(t, got, want)
	})
}
