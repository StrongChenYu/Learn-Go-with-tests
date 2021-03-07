package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test" : "chenyu"}


	t.Run("known words", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "chenyu"

		assertEquals(t, got, want)
	})


	t.Run("unknown words", func(t *testing.T) {
		_, got := dictionary.Search("test1")

		assertErrors(t, got, ErrNotFound)
	})
}

func assertEquals(t *testing.T, got,want string)  {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t *testing.T, got, want error)  {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
