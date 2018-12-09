package codility

// Returns max binary gap of the number
func maxBinaryGap(n int) int {
    var mask int = 1
    var currGap, maxGap int = 0, 0
    inGap := false
    for mask < n {
        bit := n & mask
        if (bit != 0) && inGap {
            maxGap = maxInt(currGap, maxGap)
            currGap = 0
        } else if bit != 0 {
            inGap = true
        } else if (bit == 0) && inGap {
            currGap += 1
        }
        mask <<= 1
    }
    return maxGap
}

func maxInt(a, b int) int {
    if a > b { return a }
    return b
}