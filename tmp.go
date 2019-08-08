package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 定义局部变量 */
	i2 := 20
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

	s1 := "10"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
}