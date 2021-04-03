package main

//405. 数字转换为十六进制数
import "fmt"

//定的数确保在32位有符号整数范围内
//正数补码+相反数补码 = 能表示的最大数
func toHex(num int) string {
	if num < 0 {
		return fmt.Sprintf("%x", num+1<<32)
	}
	return fmt.Sprintf("%x", num)
}
