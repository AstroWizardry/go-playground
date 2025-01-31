package controlstructures

import "fmt"

func DemonstrateControlStructures() {
    x, y := 5, 10

    if x > y {
        fmt.Println("x is greater than y")
        return
    }
    
    if x < y {
        fmt.Println("x is less than y")
        return
    }
    
    fmt.Println("x is equal to y")

    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    switch x {
    case 1:
        fmt.Println("x is 1")
    case 2:
        fmt.Println("x is 2")
    default:
        fmt.Println("x is not 1 or 2")
    }

    switch day := "Monday"; day {
    case "Monday":
        fmt.Println("Today is Monday")
    case "Tuesday":
        fmt.Println("Today is Tuesday")
    default:
        fmt.Println("Midweek pepesadge :(")
    }
}