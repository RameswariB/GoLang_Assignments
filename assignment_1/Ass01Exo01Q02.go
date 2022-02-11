/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


3 –	Make a program that calculates an employee’s gross salary for a week, given their hourly rate and the number of hours worked.


*/

package main

import "fmt"

func main() {

	var NumberOfHours, payRate, grossSalary float64

	fmt.Print("Enter NUmber of hour worked :")
	fmt.Scan(&NumberOfHours)
	fmt.Print("Enter Pay rate per hour :")
	fmt.Scan(&payRate)
	grossSalary = NumberOfHours * payRate
	fmt.Printf("Gross salary of the employee is  %.2f", grossSalary)

}
