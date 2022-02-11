package main

type myTreeNode struct {
	node *treeNode
}

func (myNode *myTreeNode) postOrder() {
	myNode.node.setValue(1)
	myNode.node.print()
}

func main() {
}
