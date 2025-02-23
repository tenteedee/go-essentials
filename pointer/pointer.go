package pointer

import "fmt"

func Main() {
	age := 25
	fmt.Println("Age:", age)
	fmt.Println("Age address:", &age)
	fmt.Println("Adult years:", getAdultYears(&age))
}

func getAdultYears(age *int) int {
	return *age - 18
}
