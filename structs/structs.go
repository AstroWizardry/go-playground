package structs

import "fmt"

type Animal struct {
    Name     string
    Category string
}

func DemonstrateStructs() {
    animal := Animal{
        Name:     "Dog",
        Category: "Mammal",
    }

    animal.Name = "Cat"

    anotherAnimal := Animal{
        Name:     "Dog",
        Category: "Mammal",
    }

    fmt.Println(animal)
    fmt.Println(anotherAnimal)
}