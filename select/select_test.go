package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type serverConfig struct {
	delay time.Duration
}

type serverOption func(*serverConfig)

func WithDelay(d time.Duration) serverOption {
	return func(cfg *serverConfig) {
		cfg.delay = d
	}
}

func makeServer(opts ...serverOption) *httptest.Server {
	cfg := serverConfig{
		delay: 0,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.delay > 0 {
			time.Sleep(cfg.delay)
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers and return url of fastest", func(t *testing.T) {
		slowServer := makeServer(WithDelay(8 * time.Millisecond))
		fastServer := makeServer()

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got,_ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return error if nothing returns in 10ms", func(t *testing.T) {
		serverA := makeServer(WithDelay(11 * time.Millisecond))
		serverB := makeServer(WithDelay(12 * time.Millisecond))

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected error")
		}
	})
}
