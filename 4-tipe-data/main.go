package main

import (
	"fmt"
	"reflect"
)

func main() {
	// string
	// var name string = "muhammad"
	name := "muhammad"

	fmt.Println(name)

	// non decimal
	var positiveNumber uint8 = 89

	negativeNumber := -1243423644

	fmt.Println(positiveNumber)
	fmt.Println(reflect.TypeOf(positiveNumber))
	fmt.Println(negativeNumber)
	fmt.Println(reflect.TypeOf(negativeNumber))

	// decimal
	// var decimalNumber float32 = 2.62
	decimalNumber := 2.62

	fmt.Println(decimalNumber)
	fmt.Println(reflect.TypeOf(decimalNumber))

	// boolean
	// var eat bool = true
	eat := true
	
	fmt.Println(eat)
	fmt.Println(reflect.TypeOf(eat))

	
}