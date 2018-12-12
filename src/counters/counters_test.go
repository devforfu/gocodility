package counters

import "testing"
import "utils"

func TestMaxCounters(t *testing.T) {
    tests := []struct{
        n int
        ops []int
        counts []int
    }{
        {1, []int{1},[]int{1}},
        {3,[]int{1, 1, 1}, []int{3, 0, 0}},
        {3,[]int{1, 2, 3}, []int{1, 1, 1}},
        {3, []int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1}},
        {3,[]int{5, 5, 5, 5, 5}, []int{0, 0, 0}},
        {5, []int{3, 4, 4, 6, 1, 4, 4}, []int{3, 2, 2, 4, 2}},
    }

    for _, test := range tests {
        n, ops, expected := test.n, test.ops, test.counts
        actual := MaxCounters(ops, n)
        if !utils.IntArraysEqual(expected, actual) {
            t.Errorf("MaxCounters(%v) -> %v != %v", ops, actual, expected)
        }
    }
}