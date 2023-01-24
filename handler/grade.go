package handler

import (
	"math"
)

func Grade(a, b, c int) (string, float64) {
	var grade string

	d := calculateGPA(a)
	e := calculateGPA(b)
	f := calculateGPA(c)

	total := ToFixed((d + e + f) / 3, 2)

	if d == 0 {
		grade = "F"
		total = 0
	} else if e == 0 {
		grade = "F"
		total = 0
	} else if f == 0 {
		grade = "F"
		total = 0
	}


	if a >= 101 {
		grade = "Invalid"
	} else if b >= 101{
		grade = "Invalid"
	} else if c >= 101{
		grade = "Invalid"
	} else if total <= 1.99 && total >= 1 {
		grade = "D"
	} else if total <= 2.99 && total >= 2 {
		grade = "C"
	} else if total <= 3.49 && total >= 3 {
		grade = "B"
	} else if total <= 3.99 && total >= 3.5 {
		grade = "A-"
	} else if total <= 4.99 && total >= 4 {
		grade = "A"
	} else if total == 5 {
		grade = "A+"
	} else if total <= 0.99 && total >= 0 {
		grade = "F"
	} else {
		grade = "Invalid"
	}

	return grade, total
}

func calculateGPA(a int) float64 {
	var GPA float64

	if a <= 32 && a >= 0 {
		GPA = 0
	}else if a <= 39 && a >= 33 {
		GPA = 1
	} else if a <= 49 && a >= 40 {
		GPA = 2
	} else if a <= 59 && a >= 50 {
		GPA = 3
	} else if a <= 69 && a >= 60 {
		GPA = 3.5
	} else if a <= 79 && a >= 70 {
		GPA = 4
	} else if a <= 100 && a >= 80 {
		GPA = 5
	} 
	return GPA
}

func ToFixed(n float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))

	return math.Round(n*scale) / scale
}
