package main

import "fmt"

func main() {
    input := []int{1, 2, 3, 4, 5}
    expected := []int{4, 5, 1, 2, 3}
    rotated := Rotate(input, 2)
    fmt.Printf("rotated=%v, expected=%v", rotated, expected)
}
