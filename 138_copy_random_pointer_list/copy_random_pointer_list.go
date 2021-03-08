package main

// Definition for a Node.
type Node struct {
     Val int
     Next *Node
     Random *Node
 }


func getCloneNode(vistedMap map[*Node]*Node, p *Node) *Node{
	if v, ok := vistedMap[p]; ok {
		return v
	} else {
		v := Node{p.Val, nil, nil}
		vistedMap[p] = &v
		return &v
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	vistedMap := make(map[*Node]*Node)
	oldNode := head
	newNode := getCloneNode(vistedMap, oldNode)
	newHead := newNode
	for oldNode != nil {
		newNode.Next = getCloneNode(vistedMap, oldNode.Next)
		newNode.Random = getCloneNode(vistedMap, oldNode.Random)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return newHead
}

func main() {
	
}
