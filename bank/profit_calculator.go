package bank

import (
	"fmt"
	"os"

	"github.com/tenteedee/go-essentials/utils"
)

const fileName string = "profit.txt"

func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = utils.GetInputFloat("Enter revenue: ")
	expenses = utils.GetInputFloat("Enter expenses: ")
	taxRate = utils.GetInputFloat("Enter tax rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Print(writeProfitToFile(ebt, profit, ratio))

}

func calculateProfit(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := profit / revenue * 100
	return ebt, profit, ratio
}

func writeProfitToFile(ebt, profit, ratio float64) string {
	output := fmt.Sprintf(
		"Earnings Before Tax: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)

	err := os.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	return output
}
