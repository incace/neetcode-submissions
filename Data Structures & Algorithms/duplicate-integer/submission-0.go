func hasDuplicate(nums []int) bool {
    check := make(map[int]struct{})

    for _, i := range nums {
        if _, exists := check[i]; exists {
            return true
        }
        check[i] = struct{}{}
    }
    return false
}