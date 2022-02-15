/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-03

Given a list of non-negative integers nums, arrange them such that they form the largest number.
Note: The result may be very large, so you need to return a string instead of an integer.
Example 1:
Input: nums = [10,2]
Output: "210"
Example 2:
Input: nums = [3,30,34,5,9]
Output: "9534330"
Example 3:
Input: nums = [1]
Output: "1"



*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	input := []int{3, 30, 34, 5, 9}

	output := findLargestNumber(input)

	fmt.Printf("The resulted output String is %s ", output)

}
func findLargestNumber(input []int) string {

	onedigit := []int{}
	twodigit := []int{}
	output := ""
	for i := 0; i < len(input); i++ {
		if input[i] < 10 {
			onedigit = append(onedigit, input[i])
		} else {
			firstNum := input[i] / 10
			secondNum := input[i] % 10
			if firstNum > secondNum {
				twodigit = append(twodigit, firstNum)
				twodigit = append(twodigit, secondNum)
			} else {
				twodigit = append(twodigit, secondNum)
				twodigit = append(twodigit, firstNum)
			}

		}
	}
	sort.Sort(sort.IntSlice(onedigit))
	sort.Sort(sort.IntSlice(twodigit))
	// fmt.Print(onedigit)
	//fmt.Print(twodigit)
	for i := len(onedigit) - 1; i >= 0; i-- {
		output += strconv.Itoa(onedigit[i])
	}
	for i := len(twodigit) - 1; i >= 0; i-- {
		output += strconv.Itoa(twodigit[i])
	}
	return output
}
