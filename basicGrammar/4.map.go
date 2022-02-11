package main

import "fmt"

func main() {
	fmt.Println("------------------创建map-----------------")
	fmt.Println("----除了slice、map、func的内建类型外，都可以作为key----")
	fmt.Println("----Struct类型不包含上述字段，也可以作为key----")
	m := map[string]string{
		"name":    "rex",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("------------------遍历map-----------------")
	fmt.Println("------------------不保证顺序-----------------")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("------------------获取map-----------------")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if courseName2, ok := m["cour"]; ok {
		fmt.Println(courseName2, ok)
	} else {
		fmt.Println("key does not exist")
	}
	fmt.Println("------------------删除map-----------------")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
