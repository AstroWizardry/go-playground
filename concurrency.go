package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	go sayHello("World") // This will run in a new goroutine
	fmt.Println("Hello from main")
	time.Sleep(1 * time.Second) // This will wait for the goroutine to finish
}
