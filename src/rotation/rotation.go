package rotation

// Rotates int array k elements right
func Rotate(a []int, k int) []int {
    n := len(a)
    if k > n { k = k % n }
    if k == 0 { return a }
    buf := make([]int, n)
    copy(buf[:k+1], a[n-k:])
    copy(buf[n-k-1:], a[:k+1])
    return buf
}