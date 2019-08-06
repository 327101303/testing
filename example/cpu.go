package main

import (
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()

	runtime.GOMAXPROCS(num*20)
	for i := 0; i < 102400; i++ {
		go func() {
			for {
				t := time.NewTimer(time.Duration(1) * time.Second)
				select {
				case <-time.After(time.Microsecond):
				}
				t.Stop()
			}
		}()
	}
	time.Sleep(time.Second * 100)
}