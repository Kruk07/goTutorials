package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	var investmentAmount, years, expectedReturnRate float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)
	// outputs information
	// fmt.Println("Future Value:", futureValue)

	// fmt.Printf(`Future Value: %.2f\n
	// Future Value (adjusted for Inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}
