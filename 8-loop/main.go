package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// with argumen
	x := 0
	for x < 5 {
		fmt.Println("Angka", x)
		x++
	}

	// without argument
	y := 0
	for {
		fmt.Println("Angka", y)

		y++
		if y == 5 {
			break
		}
	}

	// break & continue
	for z := 1; z <= 10; z++ {
		if z % 2 == 1 {
			continue
		}

		if z > 8 {
			break
		}

		fmt.Println("Angka", z)
	}

	// nested loop
	for j := 0; j < 5; j++ {
    for k := j; k < 5; k++ {
        fmt.Print(k, " ")
    }

		// for print new line
    fmt.Println()
	}

	outerLoop:
	for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
					if i == 3 {
							break outerLoop
					}
					fmt.Print("matriks [", i, "][", j, "]", "\n")
			}
	}

  // loop array
  var fruits = [4]string{"apple", "grape", "banana", "melon"}

  for i := 0; i < len(fruits); i++ {
      fmt.Printf("elemen %d : %s\n", i, fruits[i])
  }

  // loop range array
  for i, fruit := range fruits {
      fmt.Printf("elemen %d : %s\n", i, fruit)
  }

  // loop with underscore
  for _, fruit := range fruits {
      fmt.Printf("nama buah : %s\n", fruit)
  }
}
