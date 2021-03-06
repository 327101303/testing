package main

import (
    . "fmt"
    "log"
    "math/rand"
    "net/http"
    _ "net/http/pprof"
    "os"
    "runtime/debug"
    "strconv"

    "github.com/327101303/mylogger"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

var dbChan chan int
var writecount int32
var readcount int32
var Lock sync.RWMutex
var Zero = false
var cpusize int = 5
var t1 = time.NewTimer(time.Millisecond * 5)
var logger  mylogger.Logger
var monitLogger mylogger.Logger
func makeBuffer() int {
    return rand.Intn(5000000000000)
}

func write(ch chan  int){
    for i :=0;i <500;i++{
        re := makeBuffer()
        select {
        case ch <- re:
            Lock.Lock()
            atomic.AddInt32(&writecount, 1)
            Lock.Unlock()
        case <-time.After(200 * time.Millisecond):
            Println("timed out")
            goto Loop
        }
    }
Loop:
}

func read(ch chan  int){
    for i :=0;i <500;i++{
        select {
        case _, _ = <-dbChan:
            Lock.Lock()
            atomic.AddInt32(&readcount, 1)
            Lock.Unlock()
        case <-t1.C:
            Println("timed out")
            t1.Reset(200*time.Millisecond)
            goto Loop
        }
    }
Loop:
}
func cpu(cpusize int){
    num := runtime.NumCPU()
    cpunum := num / cpusize
    runtime.GOMAXPROCS(cpunum)
    for i := 0; i < 1024; i++ {
        Println("cpuset")
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
}

func writeHandler(w http.ResponseWriter,r *http.Request) {
    go write(dbChan)
}
func readHandler(w http.ResponseWriter,r *http.Request) {
    go read(dbChan)
}
func closeHandler(w http.ResponseWriter,r *http.Request) {
    close(dbChan)
}
func cpuHandler(w http.ResponseWriter,r *http.Request){
    if  len(r.URL.Query()) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }
    for k, v := range r.URL.Query() {
         if k == "cpuset" {

             value, err := strconv.Atoi(v[0])
             if err != nil {
                 Println("can't convert to int")
             }else {
                 Printf("type:%T value:%#v\n", v[0], v[0])
                 cpusize = value
             }
             cpu(cpusize)
        }else{
            Printf("参数不正确，http://hostname:port/cpu?cpuset=1")
         }
    }
}
func status(m runtime.MemStats) string{

    return Sprint("%d,%d,%d,%d\n", m.HeapSys,  m.HeapAlloc,m.HeapIdle, m.HeapReleased, )
}
func timer() {
    //timer1 := time.NewTimer(2 * time.Second)
    ticker1 := time.NewTicker(5 * time.Second)

    go func(t *time.Ticker) {
        for {
            <-t.C
            title := Sprint("get ticker", time.Now().Format("2006-01-02 15:04:05"))
            Lock.RLock()
            Printf("%s,lenchan:%d,writecount:%d,readcount:%d\n",title,len(dbChan),writecount,readcount)
            Lock.RUnlock()
            //cpu2(cpusize / 5)
        }
    }(ticker1)
}

func gcHandler(w http.ResponseWriter,r *http.Request){
    //dbChan = nil
    runtime.GC()

}

func debuggcHandler(w http.ResponseWriter,r *http.Request){
    //dbChan = nil
    debug.FreeOSMemory()

}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func mkDir(dirPath string){

    exist, err := PathExists(dirPath)
    if err != nil {
        Printf("get dir error![%v]\n", err)
        return
    }

    if exist {
        Printf("has dir![%v]\n", dirPath)
    } else {
        Printf("no dir![%v]\n", dirPath)
        // 创建文件夹
        err := os.Mkdir(dirPath, os.ModePerm)
        if err != nil {
            Printf("mkdir failed![%v]\n", err)
        } else {
            Printf("mkdir success!\n")
        }
    }

}


func logHandler(w http.ResponseWriter,r *http.Request){
    //var logger = mylogger.NewConsoleLogger("debug")
    go func() {
        for i := 0; i < 100; i++ {
            sb := "管大妈"
            logger.Debug("%s是个好捧哏", sb)
            logger.Info("info 这是一条测试的日志")
            //logger.Error("error 这是一条测试的日志")
            monitLogger.Debug("123123,1231231,123123,123123")
        }
    }()
}



func main(){
    dbChan = make(chan int,90000000)
    mkDir("/app/logs/app/")
    mkDir("/app/logs/monitor/")
    logger = mylogger.NewFileLogger("debug","default.log","/app/logs/app/")
    monitLogger = mylogger.NewFileLogger("debug","default.log","/app/logs/monitor/")
    defer logger.Close()

    http.HandleFunc("/write", writeHandler)
    http.HandleFunc("/close", closeHandler)
    http.HandleFunc("/read", readHandler)
    http.HandleFunc("/gc", gcHandler)
    http.HandleFunc("/debuggc", debuggcHandler)
    http.HandleFunc("/cpu", cpuHandler)
    http.HandleFunc("/log", logHandler)
    go timer()
    http.ListenAndServe(":8080", nil)


}