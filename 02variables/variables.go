package main

import "fmt"

const LoginToken string = "asdfjhbajdkhrfvkjyhb" // First letter, if capital, makes the constant public

func main() {
	var username string = "Barca"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isGreat bool = false
	fmt.Println(isGreat)
	fmt.Printf("Variable is of type: %T \n", isGreat)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45346532455754675
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var club = "Barcelona"
	fmt.Println(club) // lexer in go decides the vartiable type by value given

	// no var style

	numberOfPlayers := 11
	fmt.Println(numberOfPlayers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
