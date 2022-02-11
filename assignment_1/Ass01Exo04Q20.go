/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


20	- In a game, the player tosses two coins. Let’s suppose that, if the first and second coin land on heads, the player wins $10;
 if the first lands on heads and the second on tails, the player wins $5; otherwise, the player loses. We want a program that reads
  the value of the two coins (heads or tails) and determines whether the player has won. If yes, it should display the amount won.


*/

package main

import "fmt"

func main() {

	var first_coin_head bool = true
	var second_coin_head bool = false
	var winAmount int

	if (first_coin_head == true) && (second_coin_head == true) {
		winAmount = 10
	} else if (first_coin_head == true) && (second_coin_head == false) {
		winAmount = 5

	} else {
		winAmount = 0
	}
	fmt.Printf("Player win %d", winAmount)

}
