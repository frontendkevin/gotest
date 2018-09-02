package main

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unSuppoprted operation", op)

	}

}
func div(a, b int) (c float32, d int) {
	return float32(a / b), a % b

}

func find(num int, numbers ...int) {
	fmt.Println(numbers)
	found := false
	for i, v := range numbers {
		if v == num {
			fmt.Println(num, "found at index", i)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not in", numbers)
	}

}
func main() {
	if result, err := eval(2, 3, "+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	//c, d := div(8, 3)
	//	//fmt.Println(c, d)
	c, _ := div(8, 3) //当函数有多个返回值得时候 我们如果不想要那么多个 只需要用_ 来表示就行了
	fmt.Println(c)
	find(5, 1, 2, 5, 16, 8, 9)

}
