package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const word = "Go!"
const countStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigrableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigrableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, word)
}

func main() {
	sleeper := &ConfigrableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
