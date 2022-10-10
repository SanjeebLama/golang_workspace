package main

// if income is less then, 1,50,000 then no tax
// if taxable income is in the range 1,50,001-300,000 then charge 10% tax
// if taxable income is in the range 3,00,001-500,000 then charge 20% tax
// if taxable income is above 5,00,001 then charge 30% tax

import "fmt"

func main() {
	var disposableIncome float64
	var income int = 1000000

	if income < 150000 {
		disposableIncome = float64(income)
	} else if income > 150000 && income <= 300000 {
		disposableIncome = float64(income * 90 / 100)
	} else if income > 300000 && income <= 500000 {
		disposableIncome = float64(income * 80 / 100)
	} else {
		disposableIncome = float64(income * 70 / 100)
	}
	fmt.Println("Disposable Income : ", disposableIncome)

}
