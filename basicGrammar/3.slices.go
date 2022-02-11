package main

import "fmt"

func main() {
	fmt.Println("创建切片")
	var s []int //Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 16)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("复制切片")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("删除切片元素")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("从头部弹出")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)
	fmt.Println("从尾部弹出")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}

func printSlice(arr []int) {
	fmt.Printf("%v 数组长度 = %d , 数组cap = %d\n", arr, len(arr), cap(arr))
}
