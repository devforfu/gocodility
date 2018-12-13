package dna

var impactFactor = map[int32]int {'C': 1, 'D': 2}

// Find the minimal nucleotide from a range of sequence DNA
func QueryGenomicRange(s string, p []int, q []int) []int {
    return nil
}

// Counts each letter presence for each prefix of string s
func countLetters(s string) [][]int {
    prev := make([]int, 4)
    next := make([]int, 4)
    counters := [][]int { prev }
    for _, char := range s {
        copy(next, prev)
        next[impactFactor[char]] += 1
        counters = append(counters, next)
        copy(prev, next)
    }
    return counters
}