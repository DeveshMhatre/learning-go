package main

import (
	"fmt"
)

func main() {
	intSlice := []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	floatSlice := []float32{1.5, 2, 3}
	fmt.Println(sumSlice[float32](floatSlice))

	intEmptySlice := []bool{}
	fmt.Println(isEmpty(intEmptySlice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}

	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
