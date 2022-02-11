/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

14 –	15 –	Write an algorithm that simulates the withdrawal of an amount of money from an ATM.
The algorithm should ask for the amount of the current balance and the amount of the withdrawal.
If the amount of the withdrawal is greater than the balance, display an error message; otherwise, display the new balance.

*/

package main

import "fmt"

func main() {

	var current_Balance float64
	var withdrawal_Balance float64

	fmt.Println("Enter Current and Withdrawal Amount :")
	fmt.Scanf("%f %f", &current_Balance, &withdrawal_Balance)

	if current_Balance > withdrawal_Balance {
		current_Balance = current_Balance - withdrawal_Balance
		fmt.Printf("After withdrawing %.2f $ current balance is %.2f $", withdrawal_Balance, current_Balance)
	} else {
		fmt.Println("Balance is insufficient for withdraw")
	}

}
