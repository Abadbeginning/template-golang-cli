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
		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
		fmt.Println("start ==>", start, "\nout lastOccurred ==>", lastOccurred)
	}

	return maxLength
}

func main() {
	a := lenOfNonRepeatingSubStr("abcdefgabcde")
	fmt.Println(a)
}
