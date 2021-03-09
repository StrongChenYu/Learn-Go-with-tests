package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string,error)
}

type SpyStore struct {
	response string
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string,error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got canceled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}


type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
//func (s *SpyStore) Cancel()  {
//	fmt.Println("cancel invoke")
//	s.cancelled = true
//}
//
//func (s *SpyStore) assertWasCanceled()  {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Errorf("store was not canceled")
//	}
//}
//
//func (s *SpyStore) assertNotCanceled()  {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Errorf("store was canceled")
//	}
//}

type StubStore struct {
	response string
}

func (s *StubStore) Cancel() {
	panic("implement me")
}

func (s *StubStore) Fetch() string {
	return s.response
}


func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprintf(writer, data)
	}
}

func main() {
	s := "chenyu"
	for _, c := range s {
		fmt.Println(c)
	}
}