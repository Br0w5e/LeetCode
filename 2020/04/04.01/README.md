# 20 有效括号匹配
## 思路
**字典+辅助栈**
- 用一个切片c保存未匹配的字符；
- symbol用来保存右括号与左括号的对应关系，因为只有右括号出现的时候才可能匹配到左括号，从切片c中出栈；
- 判断条件：栈内有字符&&当前字符与栈顶字符匹配