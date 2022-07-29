package main

import "fmt"

func main() {
    var firstname string = "Muhammad"

    var lastname string
    lastname = "Alfathoni"
    
    fmt.Println(firstname, lastname)

    // Declaration without type data
    middlename := "Ayub"

    fmt.Println(middlename)

    // Multi variable
    var first, second, third string
    first, second, third = "satu", "dua", "tiga"

    fmt.Println(first, second, third)
}