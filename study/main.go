package main

import (
	"bufio"
	"fmt"
	"os"
)

// Struct example
type Person struct {
	Name string
	Age  int
}

// Helper function example
func add(a, b int) int {
	return a + b
}

func main() {
	// Variables and pointers
	var x int = 10
	var ptr *int = &x
	fmt.Printf("Value of x: %d, Address of x: %p\n", x, ptr)

	// Input using bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)

	// Loops and range
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Rune example
	for _, r := range "alpha" {
		fmt.Printf("Rune: %c\n", r)
	}

	// Using struct
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("Person: %+v\n", person)

	// Calling helper function
	sum := add(5, 7)
	fmt.Printf("Sum: %d\n", sum)
}
