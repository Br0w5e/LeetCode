//数组中最小的k个
//先排序再处理
func getLeastNumbers(arr []int, k int) []int {
    bubbleSort(arr, 0, len(arr))
    res := arr[:k]
    return res
}

func bubbleSort(arr []int, l, r int) []int {
    if l == r {
        return arr
    } else {
        for i := l; i < r-1; i++{
            if arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
            }
        }
        return bubbleSort(arr, l, r-1)
    }
}