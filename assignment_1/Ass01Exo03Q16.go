/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

16 –	An electricity bill is calculated by obtaining the number of days for which electricity is supplied and the number
of kilowatt hours (kWh) consumed. The client is billed at a rate of $0.50 per day and $0.30 per kWh. For a client that has
 consumed more than 200 kWh, their rate is reduced from $0.30 to $0.20 for additional kWh. We want to obtain the total amount for the bill.

*/

package main

import "fmt"

func main() {

	var number_days int
	var total_amount float64
	var consumed_kilowatt int

	fmt.Println("Enter Number of days and Killowatt you have consumed :")
	fmt.Scanf("%d %d", &number_days, &consumed_kilowatt)

	if consumed_kilowatt > 200 {

		total_amount = (float64(number_days) * float64(0.03)) + (float64(consumed_kilowatt) * float64(0.03))

	} else {
		total_amount = (float64(number_days) * float64(0.05)) + (float64(consumed_kilowatt) * float64(0.02))
	}

	fmt.Printf("Total amount for the bill is %.2f", total_amount)

}
