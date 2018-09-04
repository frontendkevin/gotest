package main

import "fmt"

func main() {
	//指针变量的类型 是 *T
	a := 24
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
	//利用指针修改value
	c := 255
	d := &c
	*d++
	fmt.Println(c)
	//向函数传递指针参数
	change(d)
	fmt.Println(c)

	//我们不要想函数传递数组指针来改变数组的内容 ，我们可以传递切片给函数，因为切片是对数组的引用 因此该该拜年切片也会改变数组的
	g := [3]int{89, 90, 91}
	modify(g[:])
	fmt.Println(g)

}
func change(val *int) {
	*val = 55

}
func modify(sls []int) {
	sls[0] = 90
}
