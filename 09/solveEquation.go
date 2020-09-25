package main

//640. 求解方程
import (
	"fmt"
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	expres := strings.Split(equation, "=")
	//等号左边
	xl, nl := getValue(expres[0])
	//等号右边
	xr, nr := getValue(expres[1])

	x := xl - xr
	n := nr - nl

	if x == 0 && n == 0 {
		return "Infinite solutions"
	} else if x == 0 && n != 0 {
		return "No solution"
	} else {
		return fmt.Sprintf("x=%d", n/x)
	}
}

func getValue(expr string) (int, int) {
	//处理负号
	expr = strings.ReplaceAll(expr, "-", "+-")
	//处理X系数
	expr = strings.ReplaceAll(expr, "-x", "-1x")

	//将加号取消
	ss := strings.Split(expr, "+")

	xVal := 0
	nVal := 0

	for _, s := range ss {
		xindex := strings.Index(s, "x")
		if xindex == 0 {
			xVal += 1
		} else if xindex > 0 {
			x, _ := strconv.Atoi(s[:xindex])
			xVal += x
		} else {
			n, _ := strconv.Atoi(s)
			nVal += n
		}
	}
	return xVal, nVal
}

func main() {
	equation := "x+5-3+x=6+x-2"
	fmt.Println(solveEquation(equation))
}
