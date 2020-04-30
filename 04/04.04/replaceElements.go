//将每个元素替换为右侧最大的元素，最后一个替换为-1
//思路：从右往左替换
func replaceElements(arr []int) []int {
    n := len(arr)
    max := arr[n-1]
    arr[n-1] = -1
    for i := n-2; i >= 0; i--{
        tmp := max
        if arr[i] > max {
            max = arr[i]
        }
        arr[i] = tmp
    }
    return arr
}