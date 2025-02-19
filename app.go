package main

import (
	"fmt"
	"reflect"
)

var i, j = 1, 2

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var c, python, java = true, false, "no"
	k := 3
	fmt.Println(reflect.TypeOf(c), " ", c)
	fmt.Println(reflect.TypeOf(python), " ", python)
	fmt.Println(reflect.TypeOf(java), " ", java)
	fmt.Println(reflect.TypeOf(k), " ", k)

	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
