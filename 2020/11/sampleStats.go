package main

//1093. 大样本统计

func sampleStats(count []int) []float64 {
	status := make([]float64, 5)

	status[0],status[1],status[2],status[3],status[4] = min(count),max(count),equ(count),mid(count),man(count)
	return status
}

func mid(a []int) float64 {
	amount := 0
	for _, v:= range a {
		amount = amount + v
	}

	p, q := 0, 0
	i:=0
	for k, v := range a {

		if i<= (amount-1)/2 && i+v > (amount-1)/2 {
			q = k
		}
		if i<= amount/2 && i+v > amount/2 {
			p = k
			break
		}
		i = i + v
	}

	return float64(p+q)/2
}

func man(a []int) float64 {
	m, n := 0, 0

	for k, v:= range a {
		if m<v {
			m = v
			n = k
		}
	}
	return float64(n)
}

func equ(a []int) float64 {
	sum, amount := 0, 0
	for k, v := range a {
		sum = sum + k * v
		amount = amount + v
	}
	return float64(sum)/float64(amount)
}

func max(a []int) float64 {
	res := 0
	for k, v:=range a {
		if v>0 {
			res = k
		}
	}
	return float64(res)
}

func min(a []int) float64 {
	res := 0
	for k, v:=range a {
		if v>0 {
			res = k
			break
		}
	}
	return float64(res)
}