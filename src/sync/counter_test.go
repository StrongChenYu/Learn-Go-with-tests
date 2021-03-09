package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		fmt.Println(counter.Value())

		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		assertCounter(t, &counter, want)
	})

	t.Run("concurrent run count from 1 to 1000", func(t *testing.T) {
		counter := Counter{}
		want := 1000

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < 1000; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, &counter, want)

	})
}

func assertCounter(t *testing.T, counter *Counter, want int)  {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, but want %d", counter.Value(),  want)
	}
}