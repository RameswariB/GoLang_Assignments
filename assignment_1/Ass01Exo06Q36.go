/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


36 –	Create four algorithms, each displaying the corresponding one of the following sequences:

a)		5	10	15	20	25	30	35	40
b)		3	5	7	9	11	13	15
c)		80	70	60	50	40	30	20
d)		1	2	6	24	120	720



*/

package main

import "fmt"

func main() {

	var number1, number2, number3, number4 int
	var init_step, final_step = 0, 8

	fmt.Println("Enter Number :")
	fmt.Scanf("%d", &number1)
	number2 = number1 - 2
	number3 = number1 * 16
	number4 = 1

	for init_step < final_step {
		fmt.Printf("%d ", number1)
		number1 = number1 + 5
		init_step++
	}

	fmt.Println("")
	number1 = number1 - 2
	init_step = 0
	for init_step < final_step {
		fmt.Printf("%d ", number2)
		number2 = number2 + 2
		init_step++
	}
	fmt.Println("")
	init_step = 0
	for init_step < final_step {
		fmt.Printf("%d ", number3)
		number3 = number3 - 10
		init_step++
	}
	fmt.Println("")
	init_step = 1
	for init_step < final_step {
		number4 = init_step * number4
		fmt.Printf("%d ", number4)
		init_step++
	}

}
