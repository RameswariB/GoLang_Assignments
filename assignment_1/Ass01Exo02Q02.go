/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


2 –	We want to determine the height of a building of n floors, knowing that the ground floor has a height of 6 meters
and that the other floors each have a height of 4 meters.

*/

package main

import (
	"fmt"
)

func main() {

	var totalheight, heightOtherfloor, floor int

	fmt.Print("Enter the totalfloor number ")
	fmt.Scan(&floor)
	heightOtherfloor = (floor - 1) * int(4)
	totalheight = heightOtherfloor + int(6)
	fmt.Printf("Total height of the building id %d", totalheight)

}
