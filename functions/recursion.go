package functions

import "fmt"

func Recursion() {
	fmt.Println(factorial(5))
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
