package main

import "fmt"

func main() {
	// operator aritmatika
	a := 4
	b := 2

	sum := a + b
	substract := a - b
	multiply := a * b
	divide := a / b

	fmt.Println(sum)
	fmt.Println(substract)
	fmt.Println(multiply)
	fmt.Println(divide)

	// operator perbandingan
	isEqual := a == b
	isNotEqual := a != b
	lessThan := a < b
	lessThanEqual := a <= b
	moreThan := a > b
	moreThanEqual := a >= b

	fmt.Println("sama dengan", isEqual)
	fmt.Println("tidak sama dengan", isNotEqual)
	fmt.Println("kurang dari", lessThan)
	fmt.Println("kurang dari sama dengan", lessThanEqual)
	fmt.Println("lebih besar dari", moreThan)
	fmt.Println("lebih besar dari sama dengan", moreThanEqual)

	// operator logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("false && true \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("false || true \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!false \t\t(%t) \n", leftReverse)
}