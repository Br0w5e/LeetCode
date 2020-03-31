//排序算法大集锦

//冒泡
//从小到大
func sortArray(nums []int) []int {
	for i := 0; i < len(nums)-1; i++{
        flag := false
		for j := len(nums)-1; j > i; j--{
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
                flag = true
			}
		}
        if flag == false {
            return nums
        }
	}
	return nums
}
//从大到小
func sortArray(nums []int) []int {
	for i := 0; i < len(nums); i++{
		for j := i; j < len(nums)-1; j++{
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}


//选择排序
func sortArray(nums []int) []int {
	for i := 0; i < len(nums); i++{
		for j := i+1; j < len(nums); j++{
			if nums[i] > nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

//优化版本
func sortArray(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        t := i
        //选择最小的, 和前面的交换
        for j := i + 1; j < len(nums); j++ {
            if nums[t] > nums[j] {
                t = j
            }
        }
        nums[t], nums[i] = nums[i], nums[t]
    }
    return nums
}


//快速排序
func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}

func quickSort(nums[]int, low int, high int) {
    if low >= high {
        return
    }
    if low < high {
        pivotpos := partition(nums, low, high)
        quickSort(nums, low, pivotpos-1)
        quickSort(nums, pivotpos+1, high)
    }
}

func partition(nums []int, low int, high int) int {
    pivot := nums[low]
    for low < high {
		//比pivot大的放在右边不变，小的进行记录并与low交换
        for low < high && nums[high] >= pivot {
            high--
        }
        nums[low] = nums[high]
		//比pivot小的的放在右边不变，小的进行记录并与low交换
        for low < high && nums[low] <= pivot {
            low++
        }
        nums[high] = nums[low]
    }
    nums[low] = pivot
    return low
}

//插入排序
func sortArray(nums []int) []int {
    for i := 0; i < len(nums); i++{
		tmp := nums[i]
		j := i
		for j > 0 && nums[j-1] > tmp {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = tmp
    }
    return nums
}

func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        //不断往前交换, 插入到已排序数组中
        for j := i; j > 0; j-- {
            if nums[j-1] > nums[j] {
                nums[j], nums[j-1] = nums[j-1], nums[j]
            }
        }
    }
    return nums
}

//系统库
func sortArray(nums []int) []int {
    sort.Ints(nums)
    return nums
}

//归并排序
//55% 9%
func sortArray(nums []int) []int {
    if len(nums) <= 1 {return nums}

    left := sortArray(nums[:len(nums)/2])
    right := sortArray(nums[len(nums)/2:])
    return merge(left, right)
}

func merge(left, right []int) []int {
    temp := make([]int, 0, len(left) + len(right))

    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            temp = append(temp, left[0])
            left = left[1:]
        } else {
            temp = append(temp, right[0])
            right = right[1:]
        }
    }

    if len(left) > 0 {
        temp = append(temp, left...)
    }

    if len(right) > 0 {
        temp = append(temp, right...)
    }

    return temp
}

//shell排序
//20% 81%
func sortArray(nums []int) []int {
    n := len(nums)
    dk := n / 2
    for dk >= 1 {
        for i := dk; i < n; i++ {
            if nums[i] < nums[i-dk] {
                temp := nums[i]
                j := i - dk
                for j >= 0 && temp < nums[j] {
                    nums[j+dk] = nums[j]
                    j -= dk
                }
                nums[j+dk] = temp
            }
        }
        dk /= 2
    }
    return nums
}

//桶排序
//限制了元素范围
//需要额外空间
//98% 5%
func sortArray9(nums []int) []int {
    s := [100001]int{}
    for i := 0; i < len(nums); i++ {
        s[nums[i] + 50000]++
    }

    idx := 0
    for i := 0; i < 100001; i++ {
        //处理重复元素
        for s[i] > 0 {
            //i - 50000 == nums[i]
            nums[idx] = i - 50000
            idx++
            s[i]--
        }
    }
    return nums
}

//基数排序 其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较
//LSD 从最低位开始 排序稳定 另外有 MSD 从最高位开始排序不稳定
//最好只是正整数, 范围不能过大
func sortArray(nums []int) []int {
    if len(nums) <= 1 {return nums}

    maxl := MaxLen(nums)
    maxValue := nums[0]
    minValue := nums[0]

    for _, v := range nums {
        if v > maxValue {
            maxValue = v
        } else if v < minValue {
            minValue = v
        }
    }

    for i := range nums {
        nums[i] -= minValue
    }
    ret := RadixCore(nums, 0, maxl)
    for i := range ret {
        ret[i] += minValue
    }
    return ret
}

func RadixCore(arr []int, digit, maxl int) []int {   //核心排序机制时间复杂度 O( d( r+n ) )
    if digit >= maxl {
        return arr                                               //排序稳定
    }

    radix := 10
    count := make([]int, radix)
    bucket := make([]int, len(arr))

    for i := 0; i < len(arr); i++ {
        count[GetDigit(arr[i], digit)]++
    }

    for i := 1;i < radix; i++ {
        count[i] += count[i-1]
    }

    for i := len(arr) - 1; i >= 0; i-- {
        d := GetDigit(arr[i], digit)
        bucket[count[d]-1] = arr[i]
        count[d]--
    }
    return RadixCore(bucket,digit + 1,maxl)
}

func GetDigit(x, d int) int {                          //获取某位上的数字
    a:=[]int{1, 10, 100, 1000, 10000, 100000, 1000000}
    return (x / a[d]) % 10
}

func MaxLen(arr []int) int {                 //获取最大位数
    var maxl, curl int
    for i := 0; i < len(arr); i++ {
        curl = len(strconv.Itoa(arr[i]))
        if curl > maxl {
            maxl = curl
        }
    }
    return maxl
}

//计数排序
func sortArray13(nums []int) []int {
    maxValue := nums[0]
    minValue := nums[0]

    for _, v := range nums {
        if v > maxValue {
            maxValue = v
        } else if v < minValue {
            minValue = v
        }
    }

    bucketLen := maxValue - minValue + 1
    bucket := make([]int, bucketLen) // 初始为0的数组

    for _, v := range nums {
        bucket[v - minValue]++
    }

    sortedIndex := 0
    for j := 0; j < bucketLen; j++ {
        for bucket[j] > 0 {
            nums[sortedIndex] = j + minValue
            sortedIndex++
            bucket[j]--
        }
    }

    return nums
}

//三路快排
//适合大量数值相等的情况
func sortArray14(nums []int) []int {
    l := 0
    r := len(nums) - 1
    if (l >= r) {return nums}
    t := nums[l + (r - l) / 2]
    tl := l
    tr := r
    for i := tl; i <= tr; {
        if nums[i] == t {
            i++
        } else if nums[i] < t {
            nums[i], nums[tl] = nums[tl], nums[i]
            i++
            tl++
        } else {
            nums[i], nums[tr] = nums[tr], nums[i]
            tr--
        }
    }
    sortArray(nums[l:tl])
    sortArray(nums[tl+1:])
    return nums
}