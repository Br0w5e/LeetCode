//两个有序数组合并为一个，第二个添加到第一个末尾
func merge(nums1 []int, m int, nums2 []int, n int)  {
    p1, p2, p := m-1, n-1, m+n-1
    for p2 >= 0 {
        if p1 >= 0 && nums1[p1] > nums2[p2]{
            nums1[p] = nums1[p1]
            p1--
        } else {
            nums1[p] = nums2[p2]
            p2--
        }
        p--
    }
    return
     
}
