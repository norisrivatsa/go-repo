package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there !")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

	var myEngine3 electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine3, 50)
}
