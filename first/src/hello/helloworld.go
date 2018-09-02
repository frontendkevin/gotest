package main

import (
	"fmt"
	"math"
)

func variable()  {
	var a int
	var s string
	var bol1,bol2 bool = true,true
	fmt.Printf("%d  %q %b %b",a,s,bol1,bol2)

}

func printVal()  {
	var a = 0
	var b="hi girl"
	//冒号的写法是不能再函数外面用的
	c := "world"
	fmt.Println(a,b,c)

}
func printConst()  {
	const filename string = "this is a string"
	const a,b int = 2, 3
	var c int
	c = int(math.Sqrt(4))
	fmt.Println(filename,a,b,c)

}
func enums()  {
	const (
		cpp = iota
		java
		golang
		javascript
	)
	fmt.Println(cpp,java,golang,javascript)

}

func main(){
	fmt.Printf("hello world")
	variable()
	printVal()
	printConst()
	enums()
}

