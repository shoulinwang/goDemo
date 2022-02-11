package main

import "fmt"

func main() {
	//寻找最长不含有重复字符的字串
	//abcabcbb-》abc
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbbacd"))
	fmt.Println(lengthOfNonRepeatingSubStr("qwerwqer"))
	fmt.Println(lengthOfNonRepeatingSubStr("ebdgasdgfas"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("1"))
	fmt.Println(lengthOfNonRepeatingSubStr("我是我的方法"))
	fmt.Println(lengthOfNonRepeatingSubStr("的撒噶高发多发第三方爱上各位发到付在v咋日期去微软他阿发的发射点"))

}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0 //当前找到的最长不含有重复字符的子串的第一个字符的下标
	maxLength := 0
	for i, ch := range []rune(s) {
		lastIndex, ok := lastOccurred[ch]
		if ok && lastIndex >= start {
			start = lastIndex + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
