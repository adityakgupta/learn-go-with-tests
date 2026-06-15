package main

import (
	"fmt"
	"net/http"
	"time"
)

func ping(s string) chan struct{} {
	c := make(chan struct{})

	go func() {
		http.Get(s)
		close(c)
	}()

	return c
}

func Racer(a, b string) (res string, err error) {
	select {
	case <-ping(a):
		res = a
	case <-ping(b):
		res = b
	case <-time.After(10 * time.Millisecond):
		return "", fmt.Errorf("timed out, waiting for %s and %s", a, b)
	}

	return
}
