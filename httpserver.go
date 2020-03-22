package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)
var filename = "txt"
func tracefile(file,str_content string) {
    fd, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    var fd_time = time.Now().Format("2006-01-02 15:04:00");
    fd_content := strings.Join([]string{ str_content, }, "")
    fmt.Printf(strings.Join([]string{fd_time,",", str_content, }, ""))
    buf := []byte(fd_content)
    fd.Write(buf)
    fd.Close()
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
    content := fmt.Sprintf("hostname=%s,value=%s\n",value1,value2)
    data := []byte(content)
    tracefile(filename,string(data))
}

func main(){
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}