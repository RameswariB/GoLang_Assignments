/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


5 –	Monique want to have a little program that allows her to evaluate the total amount of her expenses for a month,
as well as the amount of money she can allocate for her leisure activities and non-essential spending.
 The program should read her projections for expenses in the following categories:

Weekly expenses (one time per week; assuming that 1 month = 4 weeks)
•	Food expenses and household expenses
•	Common expenses

Monthly expenses (one time per month)
•	Public transit pass
•	Rent
•	Total of monthly bills

The program should also read the amount of her paycheques. Monique receives two paycheques per month.

The program should then display her total expenses, her total income, and the difference


*/

package main

import "fmt"

type WeeklyExpenses struct {
	food      int
	household int
	common    int
}
type MonthlyExpenses struct {
	transit int
	rent    int
	bill    int
}
type paycheques struct {
	amount int
}

// Creating nested structure
type TotalExpenses struct {

	// structure as a field
	firstWeek, secondWeek WeeklyExpenses
	monthly               MonthlyExpenses
	firstPay, secondPay   paycheques
	total                 int
	totalPay              int
	diff                  int
}

func main() {

	// Initializing the fields
	// of the structure
	result := TotalExpenses{

		firstWeek:  WeeklyExpenses{50, 40, 20},
		secondWeek: WeeklyExpenses{50, 40, 20},
		monthly:    MonthlyExpenses{54, 320, 70},
		firstPay:   paycheques{514},
		secondPay:  paycheques{520},
		total:      0, totalPay: 0, diff: 0,
	}

	result.total = result.firstWeek.food + result.firstWeek.household + result.firstWeek.common + result.secondWeek.food + result.secondWeek.household + result.secondWeek.common + result.monthly.transit + result.monthly.rent + result.monthly.bill
	result.totalPay = result.firstPay.amount + result.secondPay.amount
	result.diff = result.totalPay - result.total
	// Display the values
	fmt.Printf("Total pay : %d \n monthy expenses : %d \n Total Saving : %d \n", result.totalPay, result.total, result.diff)

}
