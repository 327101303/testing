package main

import (
	"fmt"
	"math/rand"
	//"net/http"
	//_ "net/http/pprof"

	"runtime"
	"sync"
	"sync/atomic"
	"time"
)
//var filename = "txt"
var dbChan chan int
var writecount int32
var readcount int32
var Lock sync.RWMutex
var Zero = false
func makeBuffer() int {
	return rand.Intn(5000000000000)
}

func write(ch chan  int){
	for i :=0;i <500;i++{
		re := makeBuffer()
		//fmt.Println(re)
		ch <- re
		Lock.Lock()
		atomic.AddInt32(&writecount,1)
		Lock.Unlock()

	}
}

//func read(ch chan  int){
//	for i :=0;i <500;i++{
//		_, ok:= <- ch
//		if ok == false{
//			break
//		}
//		Lock.Lock()
//		atomic.AddInt32(&readcount, 1)
//		Lock.Unlock()
//
//	}
//}




//func status(m runtime.MemStats) string{
//
//	return Sprint("%d,%d,%d,%d\n", m.HeapSys,  m.HeapAlloc,m.HeapIdle, m.HeapReleased, )
//}
func timer() {
	//timer1 := time.NewTimer(2 * time.Second)
	ticker1 := time.NewTicker(5 * time.Second)

	go func(t *time.Ticker) {
		for {
			<-t.C
			title := fmt.Sprint("get ticker", time.Now().Format("2006-01-02 15:04:05"))
			Lock.RLock()
			fmt.Printf("%s,lenchan:%d,writecount:%d,readcount:%d\n",title,len(dbChan),writecount,readcount)
			Lock.RUnlock()

		}
	}(ticker1)
}

//func gcHandler(w http.ResponseWriter,r *http.Request){
//	dbChan = nil
//	runtime.GC()
//
//}



func main(){
	var ms runtime.MemStats

	dbChan = make(chan int,90000000)
	write(dbChan)

	go timer()
	runtime.ReadMemStats(&ms)
	fmt.Println("after exe: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC,"\n")
	fmt.Printf("%d,%d,%d,%d\n", ms.HeapSys,  ms.HeapAlloc,ms.HeapIdle, ms.HeapReleased, )
	write(dbChan)

	time.Sleep(10 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("in exe after: ", ms, "next gc:", ms.NextGC, "gc num:", ms.NumGC,"\n")
	fmt.Printf("%d,%d,%d,%d\n", ms.HeapSys,  ms.HeapAlloc,ms.HeapIdle, ms.HeapReleased, )



}