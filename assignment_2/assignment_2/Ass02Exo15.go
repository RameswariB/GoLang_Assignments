/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-03

A message containing letters from A-Z can be encoded into numbers using the following mapping:
"A" : "1"
"B" : "2"
...
"Z" : "26"

To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:
"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into "F" since "6" is different from "06".
Given a string s containing only digits, return the different ways to decode it each as an list.

Example 1:
Input: s = "12"
2 ways.
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
Example 2:
Input: s = "226"
3 ways.
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).





*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string

	fmt.Print("Enter String :")
	fmt.Scan(&input)
	result := FindMappingNumber(strings.ToUpper(input))
	fmt.Print(result)
}
func FindMappingNumber(input string) []string {
	letters := map[string]string{
		"A": "1", "B": "2", "C": "3", "D": "4", "E": "5", "F": "6", "G": "7", "H": "8",
		"I": "9", "J": "10", "K": "11", "L": "12", "M": "13", "N": "14", "O": "15", "P": "16",
		"Q": "17", "R": "18", "S": "19", "T": "20", "U": "21", "V": "22", "W": "23", "X": "25", "Y": "25", "Z": "26"}
	var output string
	result := []string{}
	for i := 0; i < len(input); i++ {
		output = letters[string(input[i])]
		result = append(result, output)
		//fmt.Print(output)
	}
	return result
}
