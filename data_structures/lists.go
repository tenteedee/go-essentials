package data_structures

import "fmt"

type Product struct {
	Id    string
	Name  string
	Price float64
}

// func List() {
// 	hobbies := [3]string{"Reading", "Swimming", "Playing"}
// 	// 1
// 	fmt.Println(hobbies)

// 	// 2
// 	fmt.Println(hobbies[0])
// 	fmt.Println(hobbies[1:3])

// 	// 3
// 	mainHobbies := hobbies[:2]
// 	fmt.Println(mainHobbies)

// 	// 4
// 	fmt.Println(cap(mainHobbies))
// 	mainHobbies = mainHobbies[1:3] // re-slicing
// 	fmt.Println(mainHobbies)

// 	// 5
// 	courseGoals := []string{"Learn Go", "Learn Java"}
// 	fmt.Println(courseGoals)

// 	// 6
// 	courseGoals[1] = "Learn JavaScript"
// 	courseGoals = append(courseGoals, "Learn Python")
// 	fmt.Println(courseGoals)

// 	// 7
// 	products := []Product{
// 		{Id: "p1", Name: "iPhone", Price: 1.1},
// 		{Id: "p2", Name: "iPad", Price: 2.2},
// 	}
// 	fmt.Println(products)

// 	newProduct := Product{Id: "p3", Name: "MacBook", Price: 3.3}
// 	products = append(products, newProduct)
// 	fmt.Println(products)
// }

func List() {
	prices := []float64{}
	prices = append(prices, 1.1, 2.2, 3.3)
	fmt.Println(prices)

	discountPrices := []float64{4.4, 5.5}
	prices = append(prices, discountPrices...) // spread operator

	fmt.Println(prices)
}

// func List() {
// 	var productNames [5]string
// 	productNames[2] = "Product 1"
// 	fmt.Println(productNames)

// 	prices := []float64{1.1, 2.2, 3.3, 4.4}
// 	fmt.Println(prices)

// 	// featuredPrices := prices[1:3]
// 	featuredPrices := prices[:3]
// 	// featuredPrices := prices[1:]

// 	fmt.Println(featuredPrices)

// 	featuredPrices[2] = 5.5
// 	fmt.Println(prices)
// 	// a slice is a reference to an array (a pointer to an array)
// 	// changing any value of a slice is changing the value of the original array
// 	fmt.Println("Length of prices: ", len(prices))
// 	fmt.Println("Capacity of prices: ", cap(prices))
// 	fmt.Println("Length of featuredPrices: ", len(featuredPrices))
// 	fmt.Println("Capacity of featuredPrices: ", cap(featuredPrices))

// }
