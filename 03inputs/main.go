package main

import "fmt"

func main() {
	var age int
	var name string
	var funFact string

	fmt.Print("Enter your name: ")
	fmt.scanln(&name)
	fmt.Println("your name is", name)
	fmt.Print("Enter your age: ")
	fmt.scanln(&age)
	fmt.Println
}
