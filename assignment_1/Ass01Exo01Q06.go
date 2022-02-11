/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


6 –	A dealership selling new vehicles asks you to construct a program that calculates the compensation paid to their salespeople.
 The base salary for all the salespeople is $400. For each vehicle sold, the salesperson receives a commission of $200. Also,
  the salesperson receives a bonus of 5% of the total amount of all their sales.
Make the program for a salesperson.

*/

package main

import (
	"fmt"
)

type salesperson struct {
	name        string
	baseSalary  float64
	vehicleSold int
	commission  float64
	bonus       float64
}

func main() {

	saleP := salesperson{"Rameswari", 400, 3, 0, 0}
	saleP.commission = float64(saleP.vehicleSold) * float64(200)
	saleP.bonus = saleP.commission * float64(0.05)

	fmt.Printf("Sales person name %s \n is having base salary %.2f $ \n and his sold %d vehicle in this month \n and recieved comission  %.2f $\n and make bonus %.2f $from all the sales", saleP.name, saleP.baseSalary, saleP.vehicleSold, saleP.commission, saleP.bonus)

}
