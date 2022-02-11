package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//为结构定义方法
func (node treeNode) print() {
	fmt.Println(node.value, " ")
}

//使用指针最为方法的接收者
//nil指针也可以调用方法
func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *treeNode) traverseFunc(f func(node *treeNode)) {
	if node == nil {
		return
	}
	node.left.traverseFunc(f)
	node.print()
	node.right.traverseFunc(f)
}

func CreatNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = CreatNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println(root)
	fmt.Println(root.right.left)
	fmt.Println("----------------遍历结构体-------------------")
	root.right.left.setValue(4)
	root.right.left.print()
	pRoot := &root
	pRoot.setValue(200)
	pRoot.print()

	/*var root2 *treeNode
	root2.setValue(201)//会报错*/
	fmt.Println("----------------遍历结构体-------------------")
	root.traverse()
}
