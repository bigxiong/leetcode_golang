package main
type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curLevel := root
	for curLevel.Left != nil {
		p := curLevel // left most Node
		for p != nil {
			// 1. 串联我的左，右子节点
			p.Left.Next = p.Right
			// 2. 串联跨父节点情况
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
			p = p.Next
		}
		curLevel = curLevel.Left
	}

	return root
}

func main() {
	
}
