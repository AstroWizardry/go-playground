package main

import "fmt"

type Animal struct {
	Name string
	Category string
}

func main() {
	animal := Animal{
		Name: "Dog",
		Category: "Mammal",
	}

	animal.Name = "Cat"

	anotherAnimal := Animal{
		Name: "Dog",
		Category: "Mammal",
	}

	fmt.Println(animal)
	fmt.Println(anotherAnimal)
}
