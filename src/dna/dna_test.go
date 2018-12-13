package dna

import (
    "testing"
    "utils"
)

func TestQueryGenomicRange(t *testing.T) {
    tests := []struct {
        s string
        p []int
        q []int
        result []int
    }{
        {"CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}, []int{2, 4, 1}},
        {"AAA",[]int{0, 0, 0}, []int{0, 1, 2}, []int{1, 1, 1}},
    }

    for _, test := range tests {
        result := QueryGenomicRange(test.s, test.p, test.q)
        if !utils.IntArraysEqual(result, test.result) {
            t.Errorf("(actual) %v != %v (expected)", result, test.result)
        }
    }
}