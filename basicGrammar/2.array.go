package main

import (
	"fmt"
)

/**
[10]int和[20]int不是相同的类型
调用func f(arr [10]int){}会拷贝数组
*/
func main() {
	//数组的定义
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grad [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grad)

	//数组的循环
	sum := 0
	for _, v := range arr3 {
		fmt.Println(v)
		sum += v
	}
	fmt.Println(sum)

	/**
	数组传递是值类型,不会改变原来的值
	*/
	fmt.Println("++++++++++++++数组传递是值类型++++++++++++")
	printArray(arr1)
	//printArray(arr2)不行，因为类型不同
	printArray(arr3)
	//值传递，arr3的值不变
	fmt.Println(arr1, arr3)
	/**
	使用指针传递数组，会改变原来的值
	*/
	fmt.Println("++++++++++++++使用指针传递数组,不是值传递++++++++++++")
	printArrayPointer(&arr1)
	printArrayPointer(&arr3)
	fmt.Println(arr1, arr3)

	/**
	切片Slice
	*/
	fmt.Println("+++++++++++++切片++++++++++++")
	arr4 := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr4[2:6])
	fmt.Println(arr4[2:])
	fmt.Println(arr4[:6])
	fmt.Println(arr4[:])
	fmt.Println("+++++++++++++切片，非值传递++++++++++++")
	updateSlice(arr4[2:6])
	fmt.Println(arr4)
	updateSlice(arr4[:])
	fmt.Println(arr4)
	fmt.Println("+++++++++++++切片的扩展++++++++++++")
	arr5 := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("arr5=%v , len(arr5)=%d ,cap(arr5)==%d \n", arr5, len(arr5), cap(arr5))
	s1 := arr5[2:4]
	fmt.Printf("s1=%v , len(s1)=%d ,cap(s1)==%d \n", s1, len(s1), cap(s1))
	s2 := append(s1, 5)
	s3 := append(s2, 6)
	s4 := append(s3, 7)
	fmt.Printf("s2=%v , len(s2)=%d ,cap(s2)==%d \n", s2, len(s2), cap(s2))
	fmt.Println(arr5)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

}

//数组是值类型
func printArray(arr [5]int) {
	arr[0] = 1100
	for _, v := range arr {
		fmt.Println(v)
	}
}

//传递指针
func printArrayPointer(arr *[5]int) {
	arr[0] = 10086
	for _, v := range arr {
		fmt.Println(v)
	}
}

//切片
func updateSlice(arr []int) {
	arr[0] = 6969
	for _, v := range arr {
		fmt.Println(v)
	}
}
