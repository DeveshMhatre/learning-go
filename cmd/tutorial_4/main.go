package main

import (
	"fmt"
)

func main() {
	// ****
	// Array
	// ****
	var intArr [3]int32 // Fixed length, same type

	intArr[0] = 1223

	// Indexable
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])

	// Contingous in memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// var intArrTwo [4]int32 = [4]int32{1, 2, 3, 4} // Same as below
	// intArrTwo := [4]int32{1, 2, 3, 4} // Same as below
	intArrTwo := [...]int32{1, 2, 3, 4}
	fmt.Println(intArrTwo)

	// ****
	// Slice
	// ****
	// Basically a dynamic array
	var intSlice []int32 = []int32{5, 6, 7}
	fmt.Println(intSlice[0])
	fmt.Println(intSlice[0:2])
	fmt.Println(intSlice)
	fmt.Printf(
		"The length of slice is %v and capacity is %v\n",
		len(intSlice),
		cap(intSlice),
	)
	intSlice = append(intSlice, 8)
	fmt.Println(intSlice)
	fmt.Printf(
		"The length of slice is %v and capacity is %v\n",
		len(intSlice),
		cap(intSlice),
	)

	var intSlice2 []int32 = []int32{9, 10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	fmt.Printf(
		"The length of slice is %v and capacity is %v\n",
		len(intSlice3),
		cap(intSlice3),
	)

	// ****
	// Map
	// ****
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Eve": 16, "Jason": 30}
	fmt.Println(myMap2["Adam"])

	delete(myMap2, "Jason")
	age, ok := myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v: \n", age)
	} else {
		fmt.Println("Invalid name")
	}

	// ****
	// Loops
	// ****
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Value: %v\n", name, age)
	}

	for i, v := range intArrTwo {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	var j int = 0
	for j < 10 {
		fmt.Println(j)
		j = j + 1
	}
}
