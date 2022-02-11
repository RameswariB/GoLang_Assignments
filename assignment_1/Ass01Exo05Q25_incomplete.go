/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

25 –	Write an algorithm that reads two triplets day1, month1, year1, and day2, month2, year2, representing two dates,
and that determines whether the first date comes before the second.
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var date [2][3]int
	var i int
	var date1, date2 string

	for i = 0; i < 2; i++ {
		fmt.Printf("Enter %d th day :", i+1)
		fmt.Scan(&date[i][0])
		fmt.Printf("Enter %d th month :", i+1)
		fmt.Scan(&date[i][1])
		fmt.Printf("Enter %d th year :", i+1)
		fmt.Scan(&date[i][2])
	}

	date1 = strconv.Itoa(date[0][2]) + "-" + strconv.Itoa(date[0][1]) + "-" + strconv.Itoa(date[0][0])
	date2 = strconv.Itoa(date[1][2]) + "-" + strconv.Itoa(date[1][1]) + "-" + strconv.Itoa(date[1][0])
	t1, err := time.Parse("2006-01-02", date1) //0001-01-01 00:00:00 +0000 UTC
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02", date2) //0001-01-01 00:00:00 +0000 UTC
	if err != nil {
		panic(err)
	}
	if date1 > date2 {
		fmt.Printf("date first %s comes after second %s  ", t1, t2)
	} else {
		fmt.Printf("date first %s comes after second %s  ", t2, t1)
	}

}
