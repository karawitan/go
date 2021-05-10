package main

import (
    "fmt"

    "github.com/karawitan/go/greetings"
    )

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
