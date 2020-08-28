package main

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
