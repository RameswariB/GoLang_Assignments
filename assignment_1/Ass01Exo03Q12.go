/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

12 –Write an algorithm that determines the amount to pay as a tip to a server in a restaurant. The tip should be 15% when the bill is at least $1.
*/

package main

import "fmt"

func main() {

	var bill int
	fmt.Println("Enter the Bill amount : ")
	fmt.Scanf("%d", &bill)
	var tip float64
	if bill > 0 {
		tip = float64(bill) * float64(0.15)
	}
	fmt.Printf(" %.2f $ amount you have to pay as a tip to a server in a restaurant. ", tip)

}
