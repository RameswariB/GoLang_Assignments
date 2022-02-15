/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-02

Write a program that displays the price of a cinema ticket according to the age of the client and the day of the week.

It should display the price of the ticket, the amount (in $) of discount applied, the age of the client, and the day of the week (in letters).

The base price of a ticket is $9.

The day of the week is to be encoded as an integer: 1 for Monday, …, 7 for Sunday.


Table of discounts

Day of the week

Age of client	Monday	Tuesday	Wednesday	Thursday	Friday
Saturday
Sunday
Younger than 16 years, or
65 years and older	10%	30%	30%	30%	10%
16 years and older, and
Younger than 65 years	20%			20%

*/

package main

import (
	"fmt"
	"time"
)

func main() {

	var age int
	var discount float64
	var ticketPrice float64 = 9

	fmt.Print("Enter age:")
	fmt.Scan(&age)
	discount = FindDiscount(age, ticketPrice)
	dt := time.Now().Weekday()
	fmt.Printf("Today is %s ", dt.String())
	fmt.Printf("Ticket price is %.2f \n discount is  %.2f \n and final price is %.2f \n", ticketPrice, discount, ticketPrice-discount)

}

func FindDiscount(age int, price float64) float64 {
	var discount float64 = 0

	day := time.Now().Weekday()
	if day.String() == "Monday" || day.String() == "Friday" || day.String() == "Saturday" || day.String() == "Sunday" {
		if age < 16 || age > 65 {
			discount = price * float64(0.1)
		} else if (day.String() == "Monday") && (age > 16 || age < 65) {
			discount = price * float64(0.2)
		} else {
			discount = 0
		}
	} else if day.String() == "Tuesday" || day.String() == "Wednesday" || day.String() == "Thursday" {
		if age < 16 || age > 65 {
			discount = price * float64(0.3)
		} else if (day.String() == "Thursday") && (age > 16 || age < 65) {
			discount = price * float64(0.2)
		} else {
			discount = 0
		}
	} else {
		discount = 0
	}

	return discount

}
