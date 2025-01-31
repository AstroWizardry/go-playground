package api

import (
	"fmt"
	"net/http"
)

func GetGreeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}



