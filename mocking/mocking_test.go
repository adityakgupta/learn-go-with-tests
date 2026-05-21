package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
	t.Run("prints countdown", func(t *testing.T) {
		buff := &bytes.Buffer{}

		Countdown(buff, &SpyCountdown{})

		got := buff.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("sleep before print", func(t *testing.T) {
		spysleep := &SpyCountdown{}

		Countdown(spysleep, spysleep)

		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spysleep.Calls) {
			t.Errorf("got %s, want %s", spysleep.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigrableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("slept %s, should've slept %s", spyTime.durationSlept, sleepTime)
	}
}

type SpyCountdown struct {
	Calls []string
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdown) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
