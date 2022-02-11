/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


18 –	Write an algorithm that reads two integers m and n and determines whether m is a multiple of n.

*/

package main

import "fmt"

func main() {

	var m, n int
	fmt.Print("Enter the value of m and n :")
	fmt.Scanf("%d %d", &m, &n)

	if m%n == 0 {

		fmt.Println(m, " is multiple of ", n)

	} else {
		fmt.Println(m, " is not in  multiple of ", n)
	}

}
