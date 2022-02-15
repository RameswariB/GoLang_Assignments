/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-03

25 –	Write an algorithm that reads two triplets day1, month1, year1, and day2, month2, year2, representing two dates,
and that determines whether the first date comes before the second.
*/

package main

import (
	"fmt"
)

func main() {

	var board = [3][4]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	//var result [][]int
	var i int
	var start int = 0
	var end int = 3
	var outputString string
	fmt.Println(board)
	for i = 0; i < end; i++ {
		//fmt.Println([start][i])
		outputString = outputString + board[start][i]
		fmt.Printf("[%d,%d]", start, i)
	}
	start = start + 1
	for i = 1; i < end-1; i++ {
		outputString = outputString + board[i][end-1]
		fmt.Printf("[%d,%d]", i, end-1)
		//fmt.Println([i][end-1])
	}
	for i = start; i < end; i++ {
		fmt.Printf("[%d,%d]", end-1, i)
		outputString = outputString + board[end-1][i]
		//fmt.Println([end-1][i])
	}
	fmt.Printf("\n output string is %s \n", outputString)

}
