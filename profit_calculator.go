package main

import (
	"fmt"
)

func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := profit / revenue * 100

	fmt.Println("Earnings Before Tax:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Profit Ratio:", ratio)
}
