package main

import "fmt"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	diff := 100001
	for i := 0; i < len(baseCosts); i++ {
		for j := 0; j < len(toppingCosts); j++ {
			for k := 1; k <= 2; k++ {
				cost := baseCosts[i] + toppingCosts[j]*k
				if cost == target {
					return cost
				}
				if abs(cost-target) < diff {
					diff = abs(cost-target)
				}
			}
		}
	}
	return target - diff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main()  {
	baseCosts := []int{1,7}
	toppingCosts := []int{3,4}
	target := 10
	fmt.Println(closestCost(baseCosts, toppingCosts, target))
}
