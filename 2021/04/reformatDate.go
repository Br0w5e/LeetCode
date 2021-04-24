package main

//1507. 转变日期格式
import (
	"strconv"
	"strings"
)

func reformatDate(date string) string {
	dateArr := strings.Split(date, " ")
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	monthM := make(map[string]string)
	for i, month := range months {
		si := strconv.Itoa(i + 1)
		if i+1 < 10 {
			si = "0" + si
		}
		monthM[month] = si
	}
	res := dateArr[2] + "-" + monthM[dateArr[1]] + "-"
	if len(dateArr[0]) == 3 {
		res += "0" + dateArr[0][:1]
	} else {
		res += dateArr[0][:2]
	}
	return res
}
