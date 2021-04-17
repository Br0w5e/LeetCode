package main

import "time"
//1360. 日期之间隔几天

//api
func daysBetweenDates(date1 string, date2 string) int {
	a, _ := time.Parse("2006-01-02", date1)
	b, _ := time.Parse("2006-01-02", date2)
	d := b.Sub(a)

	return abs(int(d.Hours() / 24))
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
