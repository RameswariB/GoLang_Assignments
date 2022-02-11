/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


21	- Write a program that reads 3 values, determines the greatest one, and displays it.

*/

package main

import "fmt"

func main() {

	var num1, num2, num3, max int

	fmt.Print("Enter First Number :")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number :")
	fmt.Scan(&num2)
	fmt.Print("Enter Third Number :")
	fmt.Scan(&num3)
	if num1 > num2 {
		if num1 > num3 {
			max = num1
		} else {
			max = num3
		}
	} else if num2 > num1 {
		if num2 > num3 {
			max = num2
		} else {
			max = num3
		}
	} else {
		max = 0
	}
	fmt.Printf("Maximum number is %d", max)

}
