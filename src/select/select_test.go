package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("fast time test", func(t *testing.T) {
		slowServer := generateHttpDelayServer(200 * time.Microsecond)
		fastServer := generateHttpDelayServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, 1 * time.Second)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("time out of 10s test", func(t *testing.T) {
		slowServer := generateHttpDelayServer(6 * time.Second)
		fastServer := generateHttpDelayServer(7 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL, 5 * time.Second)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func generateHttpDelayServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}
