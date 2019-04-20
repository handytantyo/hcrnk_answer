package golang

func gradingStudents(grades []int32) []int32 {
	finalGrade := make([]int32, len(grades))
	indices := 0
	for _, grade := range grades {
		mod := grade % 5
		if grade < 38 || mod < 3 {
			finalGrade[indices] = grade
		} else if mod >= 3 {
			finalGrade[indices] = grade + (5 - mod)
		}
		indices++
	}
	return finalGrade

}
