package rotation

import "testing"

func TestRotateRight(t *testing.T) {
    tests := []struct {
        in []int
        out []int
        k int
    }{
        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 0},
        {[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 1},
        {[]int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}, 2},
        {[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
        {[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}, 4},
        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5},
        {[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 6},
    }
    for _, test := range tests {
        in, out, k := test.in, test.out, test.k
        rotated := RotateRight(in, k)
        if !arraysEqual(rotated, out) {
            t.Errorf("(actual) %v != %v (expected); k = %v", rotated, out, k)
        }
    }
}

func TestRotateLeft(t *testing.T) {
    tests := []struct {
        in []int
        out []int
        k int
    }{
        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 0},
        {[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}, 1},
        {[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 2},
        {[]int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}, 3},
        {[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 4},
        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5},
        {[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}, 6},
    }
    for _, test := range tests {
        in, out, k := test.in, test.out, test.k
        rotated := RotateLeft(in, k)
        if !arraysEqual(rotated, out) {
            t.Errorf("(actual) %v != %v (expected); k = %v", rotated, out, k)
        }
    }
}

func arraysEqual(a []int, b []int) bool {
    if len(a) != len(b) { return false }
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] { return false }
    }
    return true
}