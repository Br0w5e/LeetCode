package main

//1604. 警告一小时内使用相同员工卡大于等于三次的人

import (
	"fmt"
	"sort"
)

func alertNames(keyName []string, keyTime []string) []string {
	//记录每个人的打卡时间
	m := make(map[string][]string)
	for i := 0; i < len(keyName); i++ {
		m[keyName[i]] = append(m[keyName[i]], keyTime[i])
	}
	//结果
	res := make([]string, 0)
	for k, v := range m {
		if judge(v) {
			res = append(res, k)
		}
	}
	sort.Strings(res)
	return res
}


func judge(times []string) bool {
	if len(times) < 3 {
		return false
	}
	//将时间转化为数字类型
	timesInt := turnToInt(times)
	left := 0
	//滑动窗口
	for i := 1; i < len(timesInt); i++ {
		if timesInt[i]-timesInt[left] > 60 {
			left++
		}
		if i - left >= 2 {
			return true
		}
	}
	return false
}

func turnToInt(nums []string) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		h := int(nums[i][0]-'0')*10 + int(nums[i][1]-'0')
		m := int(nums[i][3]-'0')*10 + int(nums[i][4]-'0')
		res[i] = h*60+m
	}
	sort.Ints(res)
	return res
}

func main()  {
	keyName := []string{"a","a","a","a","b","b","b","b","b","b","c","c","c","c"}
	keyTime := []string{"01:35","08:43","20:49","00:01","17:44","02:50","18:48","22:27","14:12","18:00","12:38","20:40","03:59","22:24"}
	fmt.Println(alertNames(keyName, keyTime))
}