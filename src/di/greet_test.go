package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T)  {
	buffer := bytes.Buffer{}
	t.Run("test1", func(t *testing.T) {
		Greet(buffer, "chenyu")
		got := buffer.String()
		want := "Hello, chenyu"

		assertEquals(t, got, want)
	})
}

func assertEquals(t *testing.T, got, want string)  {
	t.Helper()

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}
