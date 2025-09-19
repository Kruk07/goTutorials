package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	getUserInput("Revenue: ", &revenue)
	getUserInput("Expenses: ", &expenses)
	getUserInput("Tax Rate: ", &taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInput(text string, variable *float64) {
	fmt.Print(text)
	fmt.Scan(variable)
}
