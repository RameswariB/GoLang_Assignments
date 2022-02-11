/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

22	- Write a program that reads three values and displays them in ascending order.

*/

package main

import (
	"fmt"
)

func main() {

	var num1, num2, num3, max, min, mid int
	fmt.Print("Enter first value:")
	fmt.Scan(&num1)
	fmt.Print("Enter second value:")
	fmt.Scan(&num2)
	fmt.Print("Enter third value:")
	fmt.Scan(&num3)
	if num1 > num2 {
		if num1 > num3 {
			max = num1
			if num2 > num3 {
				mid = num2
				min = num3
			} else {
				mid = num3
				min = num2
			}
		} else {
			max = num3
			if num1 > num2 {
				min = num2
				mid = num1
			}
		}
	} else if num2 > num1 {
		if num2 > num3 {
			max = num2
			if num1 > num3 {
				mid = num1
				min = num3
			} else {
				mid = num3
				min = num1
			}
		} else {
			max = num3
			if num1 > num2 {
				min = num2
				mid = num1
			} else {
				min = num1
				mid = num2
			}
		}
	} else {
		max = 0
	}
	fmt.Printf("Maximum number is %d,%d, %d", min, mid, max)
}
