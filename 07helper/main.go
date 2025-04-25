package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	name := "abdulahi"
	fmt.Println(greet(name))
}
