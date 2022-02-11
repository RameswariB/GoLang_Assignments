/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

23	- The Ministère des Finances of Québec is adopting a project aiming to reduce taxes. Develop an algorithm that calculates taxes according to the table provided below. In addition, a 2% reduction of the tax rate is granted if the person is married. Furthermore, a 0.5% reduction is granted for each child. Finally, 8% is subtracted from the tax rate for those who have newly arrived in the province. Determine the amount of tax to be paid as a function of the information provided by the user.

 	Table of basic tax rates:

Salary	Tax rate
$0.00 to $18,000.00 	10%
$18,000.01 to $32,000.00 	20%
$32,000.01 to $60,000.00 	30%
$60,000.01 and more	40%

*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var salary, tax, addTax float64
	var isMarried, newlyimmigrant string
	var NoOfChild int
	fmt.Print("Enter Salary:")
	fmt.Scan(&salary)
	tax = FindTax(salary)
	fmt.Printf("your calculated tax is %.2f \n", tax)
	fmt.Print("Are you married (y/n)")
	fmt.Scan(&isMarried)
	if strings.ToLower(isMarried) == string('y') {
		addTax = tax * float64(0.02)
	} else {
		addTax = 0
	}
	tax = tax - addTax
	fmt.Print("Enter no of child :")
	fmt.Scan(&NoOfChild)
	if NoOfChild > 0 {
		addTax = tax * (float64(0.005) * float64(NoOfChild))
	}
	tax = tax - addTax
	fmt.Print("Are you newly Immigrants (y/n):")
	fmt.Scan(&newlyimmigrant)
	if strings.ToLower(newlyimmigrant) == string('y') {
		addTax = tax * float64(0.08)
	}
	tax = tax - addTax
	fmt.Printf("your finally calculated tax is %.2f \n", tax)

}
func FindTax(salary float64) float64 {
	var tax float64

	if salary > 0 && salary <= 18000 {
		tax = salary * float64(0.1)
	} else if (salary > 18000) && (salary <= 32000) {
		tax = salary * float64(0.2)
	} else if salary > 32000 && salary <= 60000 {
		tax = salary * float64(0.3)
	} else {
		tax = salary * float64(0.4)
	}

	return tax
}
