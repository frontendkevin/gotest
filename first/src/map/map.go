package main

import "fmt"

func main() {

	mapOne := make(map[string]int)
	mapOne["age"] = 44
	mapOne["size"] = 55

	mapTwo := map[string]int{
		"number": 33,
	}
	fmt.Println(mapOne, mapTwo["number"])

	//如何知道map中是否存在某个key,ok是true 那么就存在

	value, ok := mapTwo["number"]
	if ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("number is not in this map")
	}
	//遍历map要用for range循环 有一点很重要，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。
	for key, value := range mapOne {
		fmt.Println(key, value)
	}
	//删除map˙中的元素
	delete(mapOne, "age")
	fmt.Println(mapOne)
	//map的长度
	fmt.Println(len(mapOne))
	//和切片一样 slice是引用类型
	//== 是不能比较两个MAP的
}
