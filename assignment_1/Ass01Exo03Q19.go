/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


19 –	Give an algorithm that reads three numbers (a, b, c) and determines whether any one of the numbers is equal to the sum of the other two.
 If such a number exists, display it; otherwise, display the message “No solution”.

*/

package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Println("Enter value of a,b,c :")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a+b == c {

		fmt.Println(a, " + ", b, " = ", c)

	} else if b+c == a {

		fmt.Println(b, " + ", c, " = ", a)

	} else if a+c == b {

		fmt.Println(a, " + ", c, " = ", b)
	} else {

		fmt.Println("No Solution")

	}

}
