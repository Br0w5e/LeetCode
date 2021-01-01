//将句子中所有大写字母转化为小写
//方法一：逐个遍历，遇见大写+32
//方法二：位运算
//大写变小写、小写变大写 : 字符 ^= 32;
//大写变小写、小写变小写 : 字符 |= 32;
//小写变大写、大写变大写 : 字符 &= -33;
func toLowerCase(str string) string {
    s := []byte(str)
    /*
    for i, ele := range str {
        if ele >= 'A' && ele <= 'Z'{
            s[i] += 32
        }
    }
    */
    for i := 0; i < len(s); i++{
        s[i] |= 32
    }
    return string(s)
}