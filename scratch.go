package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    n := len(a)
    b := make([]int, n)
    b[2] = 1
    fmt.Printf("a=%v, n=%v, b=%v, len(b)=%v, cap(b)=%v", a, n, b, len(b), cap(b))
}