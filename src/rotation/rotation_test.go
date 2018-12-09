package rotation

import "testing"

func TestRotate(t *testing.T) {
    tests := []struct {
        in []int
        out []int
        k int
    }{
        {[]int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}, 2},
        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 0},
        {[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5},
    }
    for _, test := range tests {
        in, out, k := test.in, test.out, test.k
        rotated := Rotate(in, k)
        if !arraysEqual(rotated, out) {
            t.Errorf("%v != %v", rotated, out)
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