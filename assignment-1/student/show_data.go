package student

import "fmt"

type ListStudent []Student

type show func(Student) int

func SelectMember(list ListStudent, dataIdx int) {
	var student Student

	for i, data := range list {
		if i == (dataIdx - 1) {
			student = data
		}
	}

	student.showData()
}

func (student Student) showData() {
	fmt.Println("================ Student Data ================")
	fmt.Println("Name     : ", student.Name)
	fmt.Println("Address  : ", student.Address)
	fmt.Println("Job 	 : ", student.Job)
	fmt.Println("Comment  : ", student.Comment)
}
