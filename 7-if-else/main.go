package main

import "fmt"

func main() {
	age := 21

	if age == 20 {
		fmt.Println("umur 20 tahun")
	} else if age > 20 {
		fmt.Println("umur lebih dari 20 tahun")
	} else if age < 20 {
		fmt.Println("umur kurang dari 20 tahun")
	} else {
		fmt.Println("belum lahir")
	}

	var point = 10000.0

	if percent := point / 100; percent >= 100 {
        fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

    // switch age {
    // case 20:
    //     fmt.Println("masih muda")
    // case 25:
    //     fmt.Println("udah tua")
    // default:
    //     fmt.Println("tua banget")
    // }

  // switch age {
  // case 20:
  //     fmt.Println("masih muda")
  // case 21,22,23,24,25:
  //     fmt.Println("udah tua")
  // default:
  //   {
  //     fmt.Println("tua banget")
  //     fmt.Println("tua ya")
  //   }
  // }

  switch {
  case age == 20:
		fmt.Println("umur 20 tahun")
  // fallthrough harus di tempatkan di case pertama jika case true lebih dari 1
  case age > 20 && age < 25:
		fmt.Println("umur antara 20 dan 25 tahun")
    fallthrough
  case age > 20:
		fmt.Println("umur lebih dari 20 tahun")
  case age < 20:
		fmt.Println("umur kurang dari 20 tahun")
  default:
    {
		fmt.Println("umur tidak cocok")
		fmt.Println("something went wrong!")
    }
  }

  // switch case inside if else
  if age > 20 {
    switch age {
    case 22:
        fmt.Println("perfect!")
    default:
        fmt.Println("nice!")
    }
} else {
    if age == 20 {
        fmt.Println("not bad")
    } else if age == 21 {
        fmt.Println("keep trying")
    } else {
        fmt.Println("you can do it")
        if age == 0 {
            fmt.Println("try harder!")
        }
    }
}
}
