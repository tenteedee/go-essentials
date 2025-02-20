package main

import (
	"fmt"
	"math"
)

func investmentCalculator() {
	const inflationRate float64 = 2.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Printf("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Printf("Enter years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
}
