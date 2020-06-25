//判断给定字母能组成的最大回文串长度
//思路：最多只能出现一次奇数
func longestPalindrome(s string) int {
	res := make([]int, 52)
	for _, letter := range s{
		if 'a' <= letter && letter <= 'z' {
			res[letter-'a']++
		} else {
			res[letter-'A'+26]++
		}
	}

	ret := 0
	for _, num := range res {
		if num % 2 == 0 {
			ret += num
		} else {
			ret += num-1
		}
	}
	if ret < len(s) {
		return ret+1
	} else {
		return ret
	}
}
//变形一
func longestPalindrome(s string) int {
    res := make([]int, 52)
    for _, letter := range s {
        if('a' <= letter && letter <= 'z') {
            res[letter-'a']++
        } else {
            res[letter-'A'+26]++
        }
    }
    ret := 0
    for _, num := range res {
		//好好理解这
        if num%2 == 1 && ret%2 == 1 {
            ret += num-1
        } else {
            ret += num
        }
    }
    return ret
}
//变形二
func longestPalindrome(s string) int {
    res := make([]int, 52)
    for _, letter := range s {
        if('a' <= letter && letter <= 'z') {
            res[letter-'a']++
        } else {
            res[letter-'A'+26]++
        }
    }
    ret := 0
    flag := 0
    for _, num := range res {
        if num%2 == 0 {
            ret += num
        } else {
            ret += num-1
            flag = 1
        }
    }
    return ret+flag
}