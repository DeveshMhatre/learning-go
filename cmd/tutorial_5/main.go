package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := []rune("résumé")
	fmt.Println(myString)

	indexed := myString[1]
	fmt.Printf("%v: %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Printf("[%v]: %v\n", i, v)
	}

	myRune := 'a'
	fmt.Printf("myRune: %v", myRune)

	strSlice := []string{"s", "u", "p", "e", "r"}

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	catStr := strBuilder.String()
	fmt.Printf("\n%v\n", catStr)
}
