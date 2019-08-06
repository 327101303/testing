// main.go

package main
import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main () {
	var id64 int64 = 99
	fmt.Println(unsafe.Sizeof(id64))
	// method 1:
	strInt64 := strconv.FormatInt(id64, 10)
	id16, _ := strconv.Atoi(strInt64)
	fmt.Println(id16)
	fmt.Println(reflect.TypeOf(id16))
	fmt.Println(unsafe.Sizeof(id16))

}