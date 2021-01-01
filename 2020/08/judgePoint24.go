package main
//回溯法
const (
	TARGET = 24
	EPSILON = 1e-6
	ADD, MULPLITY, SUBSTRACT, DIVIDE = 0, 1, 2, 3
)

func judgePoint24(nums []int) bool {
	list := []float64{}
	//转化为float类型
	for _, num := range nums {
		list = append(list, float64(num))
	}
	return solver(list)
}

func solver(list []float64) bool {
	//0个数字
	if len(list) == 0 {
		return false
	}
	//一个数字进行判断
	if len(list) == 1 {
		return abs(list[0]-TARGET) < EPSILON
	}
	// 4 3 2 1
	size := len(list)
	//  就是一个 i,j,k构成的三维坐标，其中i,j构成的坐标因为顺序会影响所以只需要去掉对角线四个就行。而不是去掉一半
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				list2 := []float64{}
				for k := 0; k < size; k++ {
					if k != i && k != j {
						list2 = append(list2, list[k])
					}
				}
				for k := 0; k < 4; k++ {
					// k<2 是因为加法和乘法满足交换律 可排除 ；
					// 为什么要且j>i呢？？因为你起码得让一类去加或者乘吧！！
					// 要不忽略nums[i]+nums[j]，要不忽略nums[j]+nums[i],如果不加且j>i,那么就一棒子打死了，所有加法乘法都不走了。
					if k < 2 && i < j {
						continue
					}
					switch k {
					case ADD:
						list2 = append(list2, list[i]+list[j])
					case MULPLITY:
						list2 = append(list2, list[i]*list[j])
					case SUBSTRACT:
						list2 = append(list2, list[i]-list[j])
					case DIVIDE:
						if abs(list[j]) < EPSILON {
							continue
						} else {
							list2 = append(list2, list[i]/list[j])
						}
					}
					if solver(list2) {
						return true
					}
					list2 = list2[:len(list2)-1]
				}
			}
		}
	}
	return false
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
