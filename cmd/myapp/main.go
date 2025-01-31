package main

import (
    "fmt"
    "go-playground/functions"
    "go-playground/structs"
    "go-playground/concurrency"
    "go-playground/api"
)

func main() {
    fmt.Println(functions.Add(1, 2))
    fmt.Println(functions.IsEven(2))
    fmt.Println(structs.Animal{Name: "Dog", Category: "Mammal"})
    concurrency.SayHello("World")
    fmt.Println(api.GetGreeting("World"))
}