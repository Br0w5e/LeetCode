//728. 自除数
package main

//规范模拟
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for left <= right {
		tmp := getNum(left)
		if judge(tmp, left) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func getNum(num int) []int {
	res := make([]int, 0)
	for num != 0 {
		res = append(res, num%10)
		num /= 10
	}
	return res
}

func judge(nums []int, val int) bool {
	for _, num := range nums {
		if num == 0 || val%num != 0 {
			return false
		}
	}
	return true
}

//优化模拟
func selfDividingNumbers2(left int, right int) []int {
	res := make([]int, 0)
	for left <= right {
		num := left
		for num != 0 {
			//排除包含0和不能整除
			if num%10 == 0 || left%(num%10) != 0 {
				break
			}
			num /= 10
		}
		//符合要求
		if num == 0 {
			res = append(res, left)
		}
		left++
	}
	return res
}
