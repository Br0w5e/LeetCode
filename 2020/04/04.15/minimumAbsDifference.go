//寻找最小绝对值差的数组
//两次遍历
func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    var min int = 1e6
    for i := 0; i < len(arr)-1; i++{
        if arr[i+1] - arr[i] < min {
            min = arr[i+1] - arr[i]
        }
    }
    res := make([][]int, 0)
    for i := 0; i < len(arr)-1; i++{
        if arr[i+1] - arr[i] == min {
            res = append(res, []int{arr[i], arr[i+1]})
        }
    }
    return res
}