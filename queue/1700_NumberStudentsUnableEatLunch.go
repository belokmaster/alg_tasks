package main

func countStudents(students []int, sandwiches []int) int {
	count := 0

	for count < len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			count = 0
		} else {
			stud := students[0]
			students = students[1:]
			students = append(students, stud)
			count++
		}
	}

	return len(students)
}
