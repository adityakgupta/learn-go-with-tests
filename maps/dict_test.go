package maps

import (
	"testing"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, key, val string) {
	t.Helper()

	got, err := d.Search(key)

	if err != nil {
		t.Fatal(err)
	}
	assertString(t, got, val)
}

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is test"}

	t.Run("known", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is test"

		assertString(t, got, want)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := dict.Search("nothing")

		assertError(t, err, ErrNotfound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		val := "this is test"

		err := dict.Add(key, val)

		assertError(t, err, nil)
		assertDefinition(t, dict, key, val)
	})

	t.Run("old", func(t *testing.T) {
		key := "test"
		val := "this is test"
		dict := Dictionary{key: val}
		err := dict.Add(key, "new")

		assertError(t, err, ErrExist)
		assertDefinition(t, dict, key, val)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("present", func(t *testing.T) {
		key := "test"
		val := "this is test"
		dict := Dictionary{key: val}

		err := dict.Update(key, "new")

		assertError(t, err, nil)
		assertDefinition(t, dict, key, "new")
	})

	t.Run("absent", func(t *testing.T) {
		key := "test"
		val := "new"
		dict := Dictionary{}

		err := dict.Update(key, val)

		assertError(t, err, ErrNotfound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("present", func(t *testing.T) {
		key := "test"
		val := "this is test"
		dict := Dictionary{key: val}

		err := dict.Delete(key)

		assertError(t, err, nil)
	})

	t.Run("absent", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("test")

		assertError(t, err, ErrNotfound)
	})
}
