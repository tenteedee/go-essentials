package functions

import (
	"fmt"
)

func Variadic() {
	fmt.Println(calcSum(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(calcSum(numbers...))
}

func calcSum(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
