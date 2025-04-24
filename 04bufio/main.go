package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.stdin)

	fmt.Print("Enter your full name")
	name, _ := reader.ReadString('/n') //includes new line character 
	fmt.Println("Yourname is ", name)
	// fmt.println(reflect.Typeof(name))
	fmt.Print("Enter your favourite number")
	fanNum, _ := reader.Readstring('/n')
	fmt.Println("Your name is ", favNum)
		fmt.Print("Enter your fun fact")
	funfact, _ := reader.ReadString('\n')
	fmt.Println("Your fun fact is ", funFact)  
fmt.Println("")
}
