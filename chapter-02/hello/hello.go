package main

import (
    "fmt"
    "seamoonpandey.com/greetings"
)


func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}

