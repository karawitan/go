package main

import (
    "fmt"

    "github.com/karawitan/go/greeting"
    )

func main() {
    // Get a greeting message and print it.
    message := greeting.Hello("Gladys")
    fmt.Println(message)
}
