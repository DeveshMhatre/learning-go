package main

import "fmt"

func main() {
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNum float32 = 10.1
	var intNum32 int32 = 10
	var result float32 = floatNum + float32(intNum32)
	fmt.Println(result)

	var myString string = "Hello World"
	fmt.Println(myString)
	fmt.Println("Length of myString: " + fmt.Sprint(len(myString)))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var myInt64 int
	fmt.Println(myInt64)

	myVar := "hello"
	fmt.Println(myVar)

	// var var1, var2, var3 int = 1, 2, 3
	var1, var2, var3 := 1, 2, 3
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	const myConst string = "yo"
	fmt.Println(myConst)
}
