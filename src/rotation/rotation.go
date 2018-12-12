package rotation

// Rotates int array k positions right
func RotateRight(arr []int, k int) []int {
    return rotate(arr, k, true)
}

// Rotates int array k positions left
func RotateLeft(arr []int, k int) []int {
    return rotate(arr, k, false)
}

func rotate(arr []int, k int, right bool) []int {
    n := len(arr)
    if k >= n { k = k % n }
    if k == 0 || k < 0 { return arr }
    //if right {
    //    arr = append(arr[n - k:], arr[:n - k]...)
    //} else {
    //    arr = append(arr[k:], arr[:k]...)
    //}
    if right { k = n - k }
    arr = append(arr[k:], arr[:k]...)
    return arr
}