package dna

var impactFactor = map[int32]int {'A': 1, 'C': 2, 'G': 3, 'T': 4}


// Find the minimal nucleotide from a range of sequence DNA
func QueryGenomicRange(s string, ps []int, qs []int) []int {
    counts := countLetters(s)
    results := []int{}
    for i := 0; i < len(ps); i++ {
        impact := computeImpact(s, counts, ps[i], qs[i])
        results = append(results, impact)
    }
    return results
}

// Counts each letter presence for each prefix of string s
func countLetters(s string) [][]int {
    prev := empty()
    counters := [][]int { empty() }
    for _, char := range s {
        next := empty()
        copy(next, prev)
        next[impactFactor[char] - 1] += 1
        counters = append(counters, next)
        copy(prev, next)
    }
    return counters
}

func empty() []int {
    return make([]int, 4)
}

func computeImpact(s string, counts [][]int, p int, q int) int {
    if p == q { return impactFactor[int32(s[p])] }
    a, b := counts[p], counts[q + 1]
    diff := diffCounts(a, b)
    for i, x := range diff {
        if x > 0 { return i + 1 }
    }
    panic("cannot get there!")
}

func diffCounts(a []int, b []int) []int {
    diff := make([]int, 4)
    for i := 0; i < 4; i++ {
        diff[i] = b[i] - a[i]
    }
    return diff
}