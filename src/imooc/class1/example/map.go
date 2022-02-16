package main

import "fmt"

// 寻找最长不含有重复字符的子串
// 对于每个字母x
//  lastOccurred[x]不存在，或者 < start -> 无需操作
//  lastOccurred[x] >= start -> 更新start
//  更新lastOccurred[x]，更新maxLength
func lenOfNonRepeatingSubStr(s string) int {
	// 初始化值
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		fmt.Println("i ==>", i, "  ch ==>", ch, "\n in lastOccurred ==>", lastOccurred)

		// 字符索引值 > start，更新start
		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}

		// 最大长度的计算 --> 索引 + 1 - start
		// start最大就是数组长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		// 当前索引值传给lastOccurred 绑定相对应的字符key
		lastOccurred[ch] = i
		fmt.Println("start ==>", start, "\nout lastOccurred ==>", lastOccurred)
	}

	return maxLength
}

func main() {
	a := lenOfNonRepeatingSubStr("abcdefgabcde")
	fmt.Println(a)
}
