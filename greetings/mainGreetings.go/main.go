package main

import (
    "fmt"
    "greetings"
)

func main() {
    message := greetings.Hello("Alice")
    fmt.Println(message)
}
