package main

import (
	"fmt"
	. "menampilkan_data/student"
	"os"
	"strconv"
)

var students ListStudent

func init() {
	students = ListStudent{Student{Person: Person{Name: "I Komang Widnyana", Address: "Wayan Gebyag"}, Job: "Mahasiswa", Comment: "karena menarik"}}
}

func main() {
	var input = os.Args
	var (
		numInput int
		_        error
	)

	for _, num := range input {
		numInput, _ = strconv.Atoi(num)
	}
	if numInput > len(students) {
		fmt.Println("The participant number you are looking for, doesn't exist")
	} else if numInput == 0 {
		fmt.Println("Please enter the participant number you are looking for first!!!")
	} else {
		SelectMember(students, numInput)
	}
}
