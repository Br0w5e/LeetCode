package main

//190. 颠倒二进制位

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		//最低位变最高位
		ret += (num & 1) << power
		//num向右移一位
		num >>= 1
		//最高位减一
		power--
	}
	return ret
}

//逐位取出来
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		//取出当前位
		res |= num & 1 << (31 - i)
		num >>= 1
	}
	return res
}
