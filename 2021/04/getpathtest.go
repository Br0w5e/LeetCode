package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main()  {
	custom, _ := os.UserHomeDir()
	strconv.Itoa()
	fmt.Println(custom)
	fmt.Println(filepath.Join(custom, "aaa"))
}
