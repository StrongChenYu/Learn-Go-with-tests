package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)


func TestCountDown(t *testing.T) {

	t.Run("test final result", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &ConfigurableSleeper{1 * time.Second}
		CountDown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test operation order", func(t *testing.T) {
		sleeper := &CountDownOperationSpy{}
		CountDown(sleeper, sleeper)

		want := []string{
			Sleep,
			Write,
			Sleep,
			Write,
			Sleep,
			Write,
			Sleep,
			Write,
		}

		if !reflect.DeepEqual(want, sleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, sleeper.Calls)
		}
	})
}
