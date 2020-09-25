package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

//71、简化路径
//API处理
func simplifyPath(path string) string {
	return filepath.Clean(path)
}

//常规方法
//栈操作
func simplifyPath2(path string) string {
	buf := strings.Split(path, "/")
	var stack []string
	for i := 0; i < len(buf); i++ {
		if buf[i] == "" || buf[i] == "." {
			continue
		}
		if buf[i] == ".." {
			//返回上级目录相当于弹出最后一个
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, buf[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath2(path))
}
