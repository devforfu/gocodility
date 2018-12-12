package counters


// Calculate the values of counters after applying all alternating operations:
// 1) increase counter by 1;
// 2) set value of all counters to current maximum.
func MaxCounters(arr []int, n int) []int {
    counts := make([]int, n)
    maxVal := 0
    lastOpIsMax := false

    for _, op := range arr {
        if 1 <= op && op <= n {
            idx := op - 1
            counts[idx] += 1
            if counts[idx] > maxVal {
                maxVal = counts[idx]
            }
            lastOpIsMax = false
        } else if op == n + 1 {
            if lastOpIsMax { continue }
            for i := range counts { counts[i] = maxVal }
            lastOpIsMax = true
        }
    }

    return counts[:]
}