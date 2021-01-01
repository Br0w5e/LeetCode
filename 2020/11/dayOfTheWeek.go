package main

//1185. 一周中的第几天

//蔡勒公式
func dayOfTheWeek(day int, month int, year int) string {
	benchmark := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	week := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if month < 3 {
		year -= 1
	}
	return week[(year + year/4 - year/100 + year/400 + benchmark[month-1] + day) % 7]
}
