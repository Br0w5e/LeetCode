func canThreePartsEqualSum(A []int) bool {
	ans := 0
	for _, num := range A {
		ans += num
	}
	if ans%3 != 0 {
		return false
	}
	res := 0
	cnt := 0
	for _, num := range A{
		res += num
		if res == ans / 3{
			res = 0
			cnt++
		}
	}
	//>= 很重要的
	if res == 0 && cnt >= 3{
		return true
	}
	return false
}

func canThreePartsEqualSum(A []int) bool {
    ans := 0
    for _, num := range A {
        ans += num
    }
    if ans%3 != 0 {
        return false
    }
    n := len(A)
    sum1 := 0
    sum2 := 0
    for i := 0; i < n; i++{
        sum1 += A[i]
        if sum1 == ans/3{
            for j := i+1; j < n-1; j++{
                sum2 += A[j]
                if sum2 == ans/3{
                    return true
                }
            }
        }
    }
    return false