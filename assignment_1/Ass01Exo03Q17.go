/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


17 –	Write an algorithm that reads an integer and determines whether it is even or odd.

*/

package main

import "fmt"

func main() {

	var number int

	fmt.Println("Enter the Number: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {

		fmt.Println("Given number ", number, " is even !!")

	} else {
		fmt.Println("Given number ", number, " is odd !!")
	}

}
