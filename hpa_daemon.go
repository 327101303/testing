package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "runtime"
    "time"
)
//var filename = "txt"
var dbChan chan int

func makeBuffer() int {
    return rand.Intn(5000000000000)
}

func write(ch chan  int){
    for i :=0;i <10;i++{
        re := makeBuffer()
        fmt.Println(re)
        ch <- re
    }
}


//func tracefile(file,str_content string) {
//    fd, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
//    var fd_time = time.Now().Format("2019-01-02 15:04:05");
//    fd_content := strings.Join([]string{ str_content, }, "")
//    fmt.Printf(strings.Join([]string{"======", fd_time, "=====", str_content, }, ""))
//    buf := []byte(fd_content)
//    fd.Write(buf)
//    fd.Close()
//}

func writeHandler(w http.ResponseWriter,r *http.Request) {
    go write(dbChan)
}



func handler(w http.ResponseWriter, r *http.Request) {

    if  len(r.URL.Query()) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }
    var value1 string
    var value2 string
    for k, v := range r.URL.Query() {
        //fmt.Printf("%s: %s\n", k, v[0])
        if k == "hostname" {
            value1 = v[0]
        }else {
            value2 = v[0]
        }
    }
    fmt.Sprintf("hostname=%s,value=%s\n",value1,value2)
    //data := []byte(content)
    //tracefile(filename,string(data))
}
//func stats(m stuct{}){
//    fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, len(dbChan), m.HeapAlloc,
//        m.HeapIdle, m.HeapReleased, )
//}
func main(){
    dbChan = make(chan int,5000000)
    http.HandleFunc("/write", writeHandler)
    http.HandleFunc("/read", handler)
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    TimerDamo := time.NewTimer(time.Duration(1)* time.Second)
    select {
    case <- TimerDamo.C:
        fmt.Println(len(dbChan))
    }


    http.ListenAndServe(":8080", nil)


}