/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


5 –	Make a program that displays the volume of a rectangular prism after reading the dimensions of the prism.


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
