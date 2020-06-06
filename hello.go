package main

import (
    "fmt"
    "time"
)
import "math/rand"

func add(a, b, c int) int {
    return a + b + c
}

func swap(str1, str2 string) (string, string) {
    return str2, str1
}

var is_go bool = true
// wrong usage
// id := 1213

func main() {
    fmt.Println(swap("world!", "Hello"))
    fmt.Println("THe current time is", time.Now())
    fmt.Println("Random number is", rand.Intn(10))
    fmt.Println("Simple add instruction: 1 + 2 + 3 =", add(1, 2, 3))

    id := 1213
    fmt.Println("Test variables is_go:", is_go, "id:", id)
}
