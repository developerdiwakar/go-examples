package main

import "fmt"

func calculateFinalGrades(grades *[]int32) {
	passGrade := 40
	for i, g := range *grades {

		graceGrade := g
		for {
			graceGrade = graceGrade + 1
			if graceGrade%5 == 0 {
				break
			}
		}
		if graceGrade-g < 3 {
			if graceGrade >= int32(passGrade) {
				(*grades)[i] = graceGrade
			}
		}
	}
}

func main() {
	// studentGrades := []int32{84, 29, 57}
	studentGrades := []int32{73, 67, 38, 33}

	calculateFinalGrades(&studentGrades)
	// Print the final grades after calculation
	fmt.Println(studentGrades)
}
