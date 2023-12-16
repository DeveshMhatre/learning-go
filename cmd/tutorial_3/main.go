package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Yo")

	division, remainder, err := intDivision(5, 0)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf(
	// 		"The division is %v",
	// 		division,
	// 	)
	// } else {
	// 	fmt.Printf(
	// 		"The division is %v and the remainder is %v",
	// 		division,
	// 		remainder,
	// 	)
	// }
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf(
			"The division is %v\n",
			division,
		)
	default:
		fmt.Printf(
			"The division is %v and the remainder is %v\n",
			division,
			remainder,
		)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var division int = numerator / denominator
	var remainder int = numerator % denominator
	return division, remainder, err
}
