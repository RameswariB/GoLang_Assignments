/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12

24 –	We want to create a program that displays the letter grade for a student, given their grade in percentage, according to the following table:

Grade	Letter
90% – 100%	A
80% – 89% 	B
70% – 79%	C
60% – 69%	D
< 60% 	F

*/

package main

import (
	"fmt"
)

type Student struct {
	grade  float64
	letter string
}

func main() {

	var studGrade float64

	fmt.Print("Enter grade (between 0 / 100):")
	fmt.Scan(&studGrade)
	stud := Student{studGrade, ""}
	stud.letter = FindLetter(stud.grade)
	fmt.Printf("Student Grade is %.2f and letter is %s", stud.grade, stud.letter)

}
func FindLetter(grade float64) string {
	var letter string
	if grade > 0 && grade < 60 {
		letter = "F"
	} else if grade >= 60 && grade < 69 {
		letter = "D"
	} else if grade >= 70 && grade < 79 {
		letter = "C"
	} else if grade >= 80 && grade < 89 {
		letter = "B"
	} else if grade >= 90 && grade < 100 {
		letter = "A"
	} else {
		letter = ""
	}

	return letter
}
