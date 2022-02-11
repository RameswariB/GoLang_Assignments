/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

11 –	Develop an algorithm that determines the greater value out of two numbers provided by the user. Display this value.
*/

package main

import "fmt"

func main() {

	var num1 int
	var num2 int
	fmt.Println("Enter First ans second Number :")
	fmt.Scanf("%d %d", &num1, &num2)

	if num1 > num2 {
		fmt.Println("Maximun Number is ", num1)
	} else {
		fmt.Println("Maximun Number is ", num2)
	}

}
