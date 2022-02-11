/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


35 –	Write an algorithm that calculates the sum of all the integers contained (inclusively) between two positive integer limits entered by the user. The program reads the smallest limit first.
Example: the sum of the integers between 5 and 10, inclusively.


*/

package main

import "fmt"

func main() {

	var number1, number2, sum int

	fmt.Println("Enter number1 and number2 :")
	fmt.Scanf("%d %d", &number1, &number2)

	for number1 <= number2 {
		sum = sum + number1
		number1 = number1 + 1
	}
	fmt.Printf("The sum of %d and %d is %d", number1, number2, sum)

}
