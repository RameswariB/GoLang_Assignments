/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


	Make a program that reads a temperature in degrees Fahrenheit and converts it into degrees Celsius. The conversion formula is:

			C=  (F-32)/9  ×5


*/

package main

import "fmt"

func main() {

	var temp_Fahrenheit, temp_celcius float64

	fmt.Print("Enter Temperatur in Fahrenheit  :")
	fmt.Scan(&temp_Fahrenheit)

	temp_celcius = ((temp_Fahrenheit - float64(32)) / 9) * 5
	fmt.Printf("Temperature in degrees Celsius  %.2f", temp_celcius)

}
