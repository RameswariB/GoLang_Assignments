/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

14 –	A print shop charges 5 cents per copy for the first 100 copies. For any subsequent copies, they charge 3 cents each.
 Write an algorithm that determines the price associated with a given number of copies.
*/

package main

import "fmt"

func main() {

	var unit int
	var sub_Price float64
	var additional_price float64
	var total_Price float64

	fmt.Println("Enter Copies (Unit) :")
	fmt.Scanf("%d", &unit)
	if unit > 100 {
		sub_Price = float64(100) * float64(0.05)
		unit = unit - 100
		additional_price = float64(unit) * float64(0.03)
		total_Price = sub_Price + additional_price
	} else {
		total_Price = float64(unit) * float64(0.05)
	}

	fmt.Printf("Total price of copy  is %.2f  ", total_Price)

}
