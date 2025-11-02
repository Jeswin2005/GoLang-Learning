package main

import (
    "fmt"
    "home/add"
    "home/sub"
)

func main() {
    // Using add package
    resultAdd := add.Sum(10, 5)
    fmt.Println("Sum:", resultAdd)

    // Using sub package
    resultSub := sub.Subtract(10, 5)
    fmt.Println("Subtraction:", resultSub)

}
