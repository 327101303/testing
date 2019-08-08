package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

//引用传递得道到是指针到拷贝，修改会同步修改结构体内到内容
func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

//值传递只是得到结构体内容到拷贝
func (b B) Print()  {
	b.Name = "BB"
	fmt.Println("B")
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}