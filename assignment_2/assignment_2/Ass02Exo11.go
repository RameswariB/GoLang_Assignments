/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-02

Given an integer array nums, return all the triplets
[nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k,    and
nums[i] + nums[j] + nums[k] == 0
Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:
Input: nums = []
Output: []
Example 3:
Input: nums = [0]
Output: []


*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}

	result := ThreeSumTriple(nums)
	fmt.Print(result)

}
func ThreeSumTriple(nums []int) [][]int {

	n := len(nums)
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				triplets := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplets[:])
				if (nums[i] + nums[j] + nums[k]) == 0 {
					tripletMap[triplets] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}
	var result [][]int
	for _, triplets := range tripletMap {
		result = append(result, triplets)
	}
	return result
}
