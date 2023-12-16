package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
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

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	var myEngineTwo electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine, 119)
	canMakeIt(myEngineTwo, 120)
}
