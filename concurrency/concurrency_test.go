package concurrency

import (
	"testing"
	"time"
)

func slowChecker(_ string) bool {
	time.Sleep(10 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "lorem ipsum"
	}

	for b.Loop() {
		CheckWebsites(slowChecker, urls)
	}
}
