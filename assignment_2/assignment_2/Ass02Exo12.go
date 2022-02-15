/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-03

Given an m x n matrix, return all elements of the matrix in spiral order.






*/

package main

import (
	"fmt"
)

func main() {

	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	result := FindSpiralMatrix(nums)
	fmt.Print(result)

}
func FindSpiralMatrix(nums [][]int) []int {

	n := len(nums)
	start := 0
	//lenght := 0
	var result []int

	for i := 0; i < n; i++ {
		result = append(result, nums[start][i])
	}
	for i := 1; i < n-1; i++ {
		result = append(result, nums[i][n-1])
	}
	for i := n - 1; i >= 0; i-- {
		result = append(result, nums[n-1][i])
	}
	for i := 0; i < n-1; i++ {
		result = append(result, nums[1][i])
	}
	return result
}
