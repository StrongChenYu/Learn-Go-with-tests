package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1,nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", url1, url2)
		
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}


//func Racer(fastUrl, slowUrl string) (winner string) {
//	aDuration := measureHttpTime(fastUrl)
//	bDuration := measureHttpTime(slowUrl)
//
//	if aDuration < bDuration {
//		return fastUrl
//	}
//	return slowUrl
//}
//
//func measureHttpTime(url string) time.Duration {
//	startA := time.Now()
//	http.Get(url)
//	return time.Since(startA)
//}