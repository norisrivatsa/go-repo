package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	fmt.Println(myString)
	var indexed = myString[1] // value will be 195
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myString1 = []rune("résumé")
	fmt.Println(myString1)
	var indexed1 = myString1[1] // value will be 233
	fmt.Printf("%v, %T\n", indexed1, indexed1)
	for i, v := range myString1 {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString1))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"b", "a", "r", "c", "e", "l", "o", "n", "a"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	var strSlice1 = []string{"b", "a", "r", "c", "e", "l", "o", "n", "a"}
	var strBuilder strings.Builder
	for i := range strSlice1 {
		strBuilder.WriteString(strSlice1[i])
	}
	var catStr1 = strBuilder.String()
	fmt.Printf("\n%v", catStr1)

}
