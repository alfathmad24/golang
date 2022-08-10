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

    // multi variable without type data
    fourth, fifth, sixth := "fourth", "fifth", "sixth"

    fmt.Println(fourth, fifth, sixth)

    // underscore variable, it uses when variable didnt used
    _ = "Muhammad"
    _ = "Alfathoni"

    name, _ := "Muhammad", "Alfathoni"

    fmt.Println(name)

    // variable with new
    age := new(string)

    fmt.Println(age)
}