package main

import "fmt"

const (
	gradeUndefine = "grade can't be define"
)

func main() {
	// ali
	fmt.Printf(`with score %d ali can get grade: %s`, 50, grade(50))
}

func grade(score int) string {
	if score <= 40 && score >= 0 {
		return "C"
	} else if score > 40 && score <= 70 {
		return "B"
	} else if score > 70 && score <= 100 {
		return "A"
	} else {
		return gradeUndefine
	}
}
