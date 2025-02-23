package bank

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func investmentCalculator() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Printf("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Printf("Enter years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)

	// fmt.Printf(`
	// Future Value: %.2f
	// Future Real Value (adjusted for inflation): %.2f`,
	// 	futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Println(formattedFV, formattedFRV)

}

func calculateFutureValues(investmentAmount float64, years float64, expectedReturnRate float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue = futureValue / math.Pow((1+inflationRate/100), years)
	// return futureValue, futureRealValue
	return
}
