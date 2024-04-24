package main

import (
	"fmt"
	"reflect"
)

func main() {

	// this is a one-line comment in Go language
	/*
		This is a multi-line comment in Go language
	*/
	// Println is a function that prints a line to the console
	fmt.Println("Hello, World!")

	// Variables

	var myFirstString string = "Hello, Go!"
	fmt.Println(myFirstString)
	myFirstString = "Hello, Go! 2"
	fmt.Println(myFirstString)

	var myFirstInt int = 42
	fmt.Println(myFirstInt)
	myFirstInt = myFirstInt + 100
	fmt.Println(myFirstInt)

	fmt.Println(myFirstString, myFirstInt)

	fmt.Println(reflect.TypeOf(myFirstString))

	var myFirstFloat float64 = 3.14
	fmt.Println(myFirstFloat)

	fmt.Println(myFirstFloat + float64(myFirstInt))

	var myFirstBool bool = true
	fmt.Println(myFirstBool)
	myFirstBool = !myFirstBool
	fmt.Println(myFirstBool)

	// Short variable declaration
	mySecondString := "Hello, Go! 3"
	fmt.Println(mySecondString)

	mySecondInt := 42
	fmt.Println(mySecondInt)

	mySecondFloat := 3.14
	fmt.Println(mySecondFloat)

	mySecondBool := true
	fmt.Println(mySecondBool)

	// Constants
	const myFirstConst string = "Hello, Go! 4"
	fmt.Println(myFirstConst)

	//	Conditional Statements
	if myFirstInt > 100 {
		fmt.Println("Greater than 100")
	} else {
		fmt.Println("Less than or equal to 100")
	}

	if myFirstInt > 100 || myFirstString == "Hello, Go! 2" {
		fmt.Println("Greater than 100 and equal to Hello, Go! 2")
	} else {
		fmt.Println("Less than or equal to 100 or not equal to Hello, Go! 2")
	}

	//	Array
	var myFirstArray [3]int = [3]int{1, 2, 3}
	fmt.Println(myFirstArray)

	//	Map
	myFirstMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(myFirstMap)

	mySecondMap := make(map[string]int)
	mySecondMap["one"] = 1
	mySecondMap["two"] = 2
	mySecondMap["three"] = 3
	fmt.Println(mySecondMap)

	//	Loops
	for i := 0; i < len(myFirstArray); i++ {
		fmt.Println(myFirstArray[i])
	}

	for key, value := range myFirstMap {
		fmt.Println(key, value)
	}

	//	Functions
	fmt.Println(add(1, 2))

	//	Structs
	type Person struct {
		name string
		age  int
	}

	myFirstPerson := Person{name: "John", age: 30}
	mySecondPerson := Person{"Jane", 25}
	fmt.Println(myFirstPerson)
	fmt.Println(mySecondPerson)
}

func add(a int, b int) int {
	return a + b
}
