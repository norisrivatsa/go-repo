package main

import "fmt"

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}
	// delete(myMap2, "Adam")
	fmt.Println(myMap2)

	for name, age := range myMap2 {
		fmt.Printf("Name : %v, Age: %v\n", name, age)
	}
}
