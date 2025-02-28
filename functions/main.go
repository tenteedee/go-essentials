package functions

import "fmt"

type transformFn func(*int) int

func Main() {
	number := []int{1, 2, 3, 4, 5}

	fmt.Println(number)

	fmt.Println(transformNumbers(&number, double))
	fmt.Println(transformNumbers(&number, triple))
	fmt.Println(transformNumbers(&number, quadruple))

	fmt.Println(transformNumbers(&number, func(value *int) int {
		return *value * 5
	}))
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	for index, value := range *numbers {
		(*numbers)[index] = transform(&value)
	}
	return *numbers
}

func double(value *int) int {
	return *value * 2
}

func triple(value *int) int {
	return *value * 3
}

func quadruple(value *int) int {
	return *value * 4
}
