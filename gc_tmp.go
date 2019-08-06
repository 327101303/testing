package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("before exe: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC)
	go exe()
	go exe()
	go exe()
	runtime.ReadMemStats(&ms)
	fmt.Println("after exe: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC)
	time.Sleep(10 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("before return: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC)
}

func exe() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("in exe before: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC)
	var mq = make(chan int64, 1000000)
	for i := 0; i < 1000000; i++ {
		mq <- int64(i)
	}
	for i := 0; i < 1000000; i++ {
		<-mq
	}
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("in exe after: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC)
	fmt.Println("len-mq:",len(mq))
}