/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


4 –	In a computer technology course, the following evaluation weights are used:

•	Laboratory work counts for 40% of the grade
•	The midterm exam counts for 25%
•	The final exam counts for 35%

Make a program that calculates the final grade of a student, assuming that each of the three grades the user inputs is out of 100


*/

package main

import (
	"fmt"
)

type Student struct {
	StudentName string
	courseName  string
	labGrade    float64
	midterm     float64
	finalGrade  float64
	Total       float64
}

func main() {

	var name, course string
	var lab, mid, final float64

	fmt.Print("Enter Student Name :")
	fmt.Scan(&name)
	fmt.Print("Enter Course Name :")
	fmt.Scan(&course)
	fmt.Print("Enter Lab Grade(out of 100) :")
	fmt.Scan(&lab)
	lab = lab * float64(0.4)
	fmt.Print("Enter Midterm Grade(out of 100) :")
	fmt.Scan(&mid)
	mid = mid * float64(0.25)
	fmt.Print("Enter Final Grade(out of 100) :")
	fmt.Scan(&final)
	final = final * float64(0.35)
	stud := Student{name, course, lab, mid, final, 0}
	stud.Total = lab + mid + final
	fmt.Printf("Student name : %s \n course Name : %s \n Lab Grade : %.1f \n Midterm Grade : %.1f \n Final Exam Grade : %.1f \n Total Grade : %.1f",
		stud.StudentName, stud.courseName, stud.labGrade, stud.midterm, stud.finalGrade, stud.Total)

}
