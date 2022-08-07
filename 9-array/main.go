package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
  // names[4] = "ez" // this line will show error because it more than the length of allocation array

	fmt.Println(names[0], names[1], names[2], names[3])

  var fruits = [4]string{"apple", "grape", "banana", "melon"}

  fmt.Println("Jumlah element \t\t", len(fruits))
  fmt.Println("Isi semua element \t", fruits)

  // var fruits [4]string

  // // cara horizontal
  // fruits  = [4]string{"apple", "grape", "banana", "melon"}

  // // cara vertikal
  // fruits  = [4]string{
  //     "apple",
  //     "grape",
  //     "banana",
  //     "melon",
  // }

  // array without length element
  var numbers = [...]int{2, 3, 2, 4, 3}

  fmt.Println("data array \t:", numbers)
  fmt.Println("jumlah elemen \t:", len(numbers))

  // array dimension
  var numbers1 = [3][4]int{{3, 2, 3}, {3, 4, 5}}
  var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

  fmt.Println("numbers1", numbers1)
  fmt.Println("numbers2", numbers2)

  // array with make
  var animals = make([]string, 2)
  animals[0] = "cat"
  animals[1] = "dog"

  fmt.Println(animals)  // [apple manggo]
}
