package main

import "fmt"

//299. 猜数字游戏

func getHint(secret string, guess string) string {
	nums := make([]int, 10)
	bulls, cows := 0, 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			//证明之前已经在guess串中出现了，guess--
			if nums[int(secret[i]-'0')] < 0 {
				cows++
			}
			//减少出现的次数（已经使用了）
			nums[int(secret[i]-'0')]++
			//证明以前已经在secret中出现了, secret++
			if nums[int(guess[i]-'0')] > 0 {
				cows++

			}
			//减少出现的次数（已经使用了）
			nums[int(guess[i]-'0')]--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func main()  {
	secret := "1807"
	guess := "7810"
	fmt.Println(getHint(secret, guess))
}