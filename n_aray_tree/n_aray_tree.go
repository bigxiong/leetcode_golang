package main


type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	var f func(*Node)
	f = func(root *Node) {
		if root != nil {
			res = append(res, root.Val)
			for _, child := range root.Children {
				f(child)
			}
		}
	}
	f(root)
	return res
}



func main() {
	
}
