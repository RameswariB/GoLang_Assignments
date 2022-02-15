/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz,
Isabela Pamplona Castilho Lima.
Date: 2022-02-03

Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:
Input: s = "  hello world  "
Output: "world hello"




*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	input := " Sky is blue "
	split := strings.Split(strings.Trim(input, " "), " ")
	fmt.Println(split)
	for i := len(split) - 1; i >= 0; i-- {
		fmt.Print(split[i] + " ")
	}
}
