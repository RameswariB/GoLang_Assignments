/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


3 –	An aircraft pilot wants to know the atmospheric pressure, expressed in atmosphere units (atm), as the weather station only
provides pressure data in kilopascal units (kPa).
1 atm is equivalent to 101.325 kPa. Make a program that performs the conversion.


*/

package main

import (
	"fmt"
)

func main() {

	var atm, kpa float64

	fmt.Print("Please enter pressure data in kilopascal units :")
	fmt.Scanf("%f", &kpa)
	atm = conversion(kpa)

	fmt.Printf("Atmospheric pressure of aircraft is %.3f ", atm)

}
func conversion(kpa float64) float64 {

	return float64(101.325) * kpa

}
