package main

import (
	"time"
)


func main() { //Main terminates before the anonymous functions even execute. Can be fixed with sleep.
	go func() {
	    for i := 0; i < 100; i++ {
		println("hello")
	    }
	}()

	go func() {
	    for i := 0; i < 100; i++ {
		println("go")
            }
	}()

	dur, _ := time.ParseDuration("1ms")
	time.Sleep(dur)
	println("ending main")
}

