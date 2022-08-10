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
  fmt.Println(animals) // ["cat", "dog", "tiger", "lion"]

  var aAnimals = animals[0:2]
  var bAnimals = animals[2:4]

  fmt.Println(aAnimals) // ["cat", "dog"]
  fmt.Println(bAnimals) // ["tiger", "lion"]

  bAnimals[0] = "cow"

  fmt.Println(animals) // ["cat", "dog", "tiger", "lion"]
  fmt.Println(bAnimals) // // ["cow", "lion"]

  // len function
  fmt.Println(len(fruits)) // ["cat", "dog", "cow", "lion"]

  // cap function
  fmt.Println(len(fruits))  // len: 4
  fmt.Println(cap(fruits))  // cap: 4 -> start from 0 until before x

  var aFruits = fruits[0:3]
  fmt.Println(len(aFruits)) // len: 3
  fmt.Println(cap(aFruits)) // cap: 4 -> start from 0 until before x

  var bFruits = fruits[1:4]
  fmt.Println(len(bFruits)) // len: 3
  fmt.Println(cap(bFruits)) // cap: 3 -> start from x until before x

  // append function
  // if len(x) === cap(x), the new variable will make new reference, if len(x) < cap(x)
  // the new variable will have the same referance
  var cFruits = fruits[0:4]

  fmt.Println(cap(cFruits))
  fmt.Println(len(cFruits))

  fmt.Println(fruits) // ["apple", "grape", "banana", "melon"]
  fmt.Println(cFruits) // ["apple", "grape"]

  var dFruits = append(cFruits, "pear")

  fmt.Println(fruits) // [apple grape pear melon]
  fmt.Println(cFruits) // ["apple", "grape", "banana", "melon", "papaya"]
  fmt.Println(dFruits)  // ["apple", "grape", "banana", "melon"]

  dFruits[0] = "watermelon"
  fmt.Println(fruits) // [apple grape pear melon]
  fmt.Println(dFruits)  // ["apple", "grape", "banana", "melon"]

  // copy function
  dst := make([]string, 3)
  src := []string{"watermelon", "pinnaple", "apple", "orange"}
  n := copy(dst, src)

  fmt.Println(dst) // watermelon pinnaple apple
  fmt.Println(src) // watermelon pinnaple apple orange
  fmt.Println(n)   // 3

  dst2 := []string{"potato", "potato", "potato"}
  src2 := []string{"watermelon", "pinnaple"}
  n2 := copy(dst2, src2)

  fmt.Println(dst2) // watermelon pinnaple potato
  fmt.Println(src2) // watermelon pinnaple
  fmt.Println(n2)   // 2

  // slices 3 index
  var fruits2 = []string{"apple", "grape", "banana"}
  var aFruits2 = fruits2[0:2]
  var bFruits2 = fruits2[0:2:2]

  fmt.Println(fruits2)      // ["apple", "grape", "banana"]
  fmt.Println(len(fruits2)) // len: 3
  fmt.Println(cap(fruits2)) // cap: 3

  fmt.Println(aFruits2)      // ["apple", "grape"]
  fmt.Println(len(aFruits2)) // len: 2
  fmt.Println(cap(aFruits2)) // cap: 3

  fmt.Println(bFruits2)      // ["apple", "grape"]
  fmt.Println(len(bFruits2)) // len: 2
  fmt.Println(cap(bFruits2)) // cap: 2
}
