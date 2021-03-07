package concurrency

import (
	"reflect"
	"testing"
	"time"
)



func mockWebSiteChecker(url string) bool  {
	if url == "www" {
		return false
	}
	return true
}

func TestCheckWebSites(t *testing.T) {
	websites := []string{
		"http://www.baidu.com",
		"http://www.bilibili.com",
		"www",
	}

	actualResult := CheckWebSites(mockWebSiteChecker, websites)

	want := len(websites)
	got := len(actualResult)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResult := map[string]bool {
		"http://www.baidu.com" : true,
		"http://www.bilibili.com": true,
		"www" : false,
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("Wanted %v, got %v", expectedResult, actualResult)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebSites(slowStubWebsiteChecker, urls)
	}
}