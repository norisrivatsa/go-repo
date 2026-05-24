package main

import "fmt"

func main() {
	var intArr [3]int32 // [length]element_types, by default [0,0,0]
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [...]int32{1, 2, 3}
	fmt.Println(intArr3)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice)) // capacity increases to 6, so array was originally len 3, with 4,5,6, now a new array with length 4 and cap 6 comes with 4,5,6,7,*,*. Cannot access the array values at indexes 4,5

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8) // 3,  8 are len and cap.
	fmt.Println(intSlice3)
	fmt.Printf("The length is %v with capacity %v", len(intSlice3), cap(intSlice3))

}
