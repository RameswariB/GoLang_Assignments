/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

13 –	A computer store sells diskettes at a price of $1 each for small orders. The store sells them at a price of
70 cents each for orders of 25+ units. Furthermore, if you are a member of Club Z, the store will give you an extra discount of 2%.
Write an algorithm that determines the total price for a purchase.
*/

package main

import "fmt"

func main() {

	var diskettes_price float64
	var unit int
	var total_Price float64
	var discount_price float64
	fmt.Println("Enter diskettes price & unit :")
	fmt.Scanf("%f %d", &diskettes_price, &unit)

	membership := true
	if unit > 25 {
		diskettes_price = float64(diskettes_price) * float64(0.70)
		if membership == true {
			discount_price = float64(diskettes_price) * float64(0.02)
		}
	}
	total_Price = diskettes_price - discount_price
	fmt.Printf("Total price of diskettes is %.2f  ", total_Price)

}
