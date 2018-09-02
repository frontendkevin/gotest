package main

import "fmt"

func main() {
	grade(80)
	fmt.Println(getScore(2))
}

func grade(score int) {
	switch {
	case score < 60:
		fmt.Println("A")
		break
	case score >= 60 && score <= 80:
		fmt.Println("S")
		break
	default:
		fmt.Println("hello")
		break
	}
}
func getScore(score int) string {
	switch score {
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	default:
		return "D"

	}
}
