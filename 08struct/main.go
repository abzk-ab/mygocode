package main

import "fmt"

// Define a struct for a student
type student struct {
	name   string
	school string
	year   int
}

func main() {
	// create a new student
	student1 := student{
		name:   "Alex",
		school: "Founders High school",
		year:   2024,
	}
	// Print student information
	fmt.Println("student information")
	fmt.Println("Name:", student1.name)
	fmt.Println("school:", student1.school)
	fmt.Println("year:", student1.year)
}
