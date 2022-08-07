package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

  // var fruitsA = []string{"apple", "grape"}      // slice
  // var fruitsB = [2]string{"banana", "melon"}    // array
  // var fruitsC = [...]string{"papaya", "grape"}  // array

  // this variable is becames a slice
  var newFruits = fruits[0:2]
  fmt.Println(newFruits) // ["apple", "grape"]

  var animals = [4]string{"cat", "dog", "tiger", "lion"}
  fmt.Println(animals)

  var aAnimals = animals[0:2]
  var bAnimals = animals[2:4]

  fmt.Println(aAnimals)
  fmt.Println(bAnimals)

  bAnimals[0] = "cow"

  fmt.Println(animals)
  fmt.Println(bAnimals)

  // len function
  fmt.Println(len(fruits)) // 4

  // cap function
  fmt.Println(len(fruits))  // len: 4
  fmt.Println(cap(fruits))  // cap: 4

  var aFruits = fruits[0:3]
  fmt.Println(len(aFruits)) // len: 3
  fmt.Println(cap(aFruits)) // cap: 4

  var bFruits = fruits[1:4]
  fmt.Println(len(bFruits)) // len: 3
  fmt.Println(cap(bFruits)) // cap: 3
}
