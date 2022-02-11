package main

import (
	"fmt"
	"unicode/utf8"
)

/**
rune相当于go的char
1、使用range遍历pos、rune对
2、使用utf8.RuneCountInString()获得字符数量
3、使用len获取字节长度
4、使用[]byte获取全部字节
*/
func main() {
	s := "YES我爱我家"
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
		//59 45 53  英文是一字节  E6 88 91   E7 88 B1  E6 88 91    E5 AE B6 中文是三字节
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d,%X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Print(ch)
		bytes = bytes[size:]
		fmt.Printf("%c\n", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d.%c)", i, ch)
	}
}
