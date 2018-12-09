package main

import "testing"

// Tests binary gap function returns correct value of max gap
func TestBinaryGap(t *testing.T) {
    tests := []struct {
        val int
        gap int
    }{
        {9,2},
        {16, 0},
        {17, 3},
        {529, 4},
    }
    for _, test := range tests {
        gap := maxBinaryGap(test.val)
        if  gap != test.gap {
            t.Errorf("%v != %v", gap, test.gap)
        }
    }
}