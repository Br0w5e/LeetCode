//猜数字
//二分法
func guessNumber(n int) int {
    left, right := 1, n
    for left < right {
        mid := (left + right)/2
        if guess(mid) == 0 {
            return mid
        } else if guess(mid) > 0 {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return left
}