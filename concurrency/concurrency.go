package concurrency

import (
	"fmt"
	"time"
)

func SayHello(name string) {
	go func() {
		fmt.Println("Hello, " + name)
	}()
	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
}