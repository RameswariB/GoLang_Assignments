/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-28

8 –	A palindrome is a number, a word, or a sentence that remains identical whether read from left to right or from right to left. For example, each of the following five-digit numbers is a palindrome: 12321, 55555, 45554, 11611. Write a program capable of reading a positive integer (greater than 0) with five digits, and of determining whether this integer is a palindrome. (Hint: use the modulus and division operators to separate the different digits composing the numbers.)
*/

package main

import "fmt"

func main() {

	var value, reverseVal string
	var isPalindrom bool = true

	fmt.Print("Enter Number or String :")
	fmt.Scan(&value)
	reverseVal = Reverse(value)
	if value[0] == '-' {
		fmt.Print("Please enter valid positive number")
		return
	} else if len(value) < 5 {
		fmt.Print("Lenght should be more than five digits")
		return
	} else {
		for i := 0; i < len(value)-1; i++ {

			for j := 0; j < len(reverseVal)-1; j++ {
				if value[i] != reverseVal[i] {
					fmt.Print(value[i])
					isPalindrom = false
					break
				}
			}
		}
		if isPalindrom == true {
			fmt.Print("The String is palindrom")
		} else {
			fmt.Print("The String is not palindrom")
		}
	}

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
