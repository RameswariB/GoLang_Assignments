/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


6 –	A dealership selling new vehicles asks you to construct a program that calculates the compensation paid to their salespeople.
 The base salary for all the salespeople is $400. For each vehicle sold, the salesperson receives a commission of $200. Also,
  the salesperson receives a bonus of 5% of the total amount of all their sales.
Make the program for a salesperson.

*/

package main

import (
	"fmt"
)

func main() {

	var height, lenght, width, volume int

	fmt.Print("Enter Height of rectagle:")
	fmt.Scan(&height)
	fmt.Print("Enter Width of rectagle:")
	fmt.Scan(&width)
	fmt.Print("Enter Lenght of rectagle:")
	fmt.Scan(&lenght)
	volume = height * width * lenght
	fmt.Printf("volume of a rectangular prism is %d  cubic units", volume)

}
