package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const firstWord = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s SpySleeper) Sleep()  {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration  time.Duration
}

func (c *ConfigurableSleeper) Sleep()  {
	time.Sleep(c.duration)
}


func CountDown(writer io.Writer, sleeper Sleeper)  {
	for i:=firstWord; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

type CountDownOperationSpy struct {
	Calls []string
}

func (spy *CountDownOperationSpy) Sleep()  {
	spy.Calls = append(spy.Calls, Sleep)
}

func (spy *CountDownOperationSpy) Write(p []byte) (n int, err error)  {
	spy.Calls = append(spy.Calls, Write)
	return 0, err
}

const Write = "write"
const Sleep = "sleep"

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	CountDown(os.Stdout, sleeper)
}
