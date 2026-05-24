package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || error ok syntax

	input, _ := reader.ReadString('\n') // err is like a catch \we can put _ if we do not care to use the error
	fmt.Println("Thanks for rating,  ", input)
	fmt.Printf("Type of this rating is %T", input) // Will be string

}
