package main

import (
	"errors"
	"fmt"
	"os"
)

const writeFile = "calculations.txt"

func main() {
	var revenue, expenses, taxRate float64

	errRevenue := getUserInput("Revenue: ", &revenue)
	if errRevenue != nil {
		panic(errRevenue)
	}

	errExpenses := getUserInput("Expenses: ", &expenses)
	if errExpenses != nil {
		panic(errExpenses)
	}

	errTaxRate := getUserInput("Tax Rate: ", &taxRate)
	if errTaxRate != nil {
		panic(errTaxRate)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	storeResults(ebt, profit, ratio)
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

func getUserInput(text string, variable *float64) error {
	fmt.Print(text)
	fmt.Scan(variable)

	if *variable < 0.0 {
		return errors.New(text + "You've entered negative number. ")
	}

	if *variable == 0.0 {
		return errors.New(text + "You've entered 0.")
	}

	return nil
}

func storeResults(ebt, profit, ratio float64) {
	text := fmt.Sprintf("Ebt: %.1f Profit: %.1f Ratio: %.3f", ebt, profit, ratio)
	os.WriteFile(writeFile, []byte(text), 0644)
}
