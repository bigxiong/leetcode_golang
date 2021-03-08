package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


type Codec struct {
	nodesVal []string
}

func Constructor() Codec {
	return Codec{nodesVal:nil}
}

// Serializes a tree to a single string.
func (this *Codec) doSerialize(root *TreeNode) {
	if root == nil {
	    this.nodesVal = append(this.nodesVal, "#")
		return
	}
	this.nodesVal = append(this.nodesVal, fmt.Sprintf("%d", root.Val))
	this.doSerialize(root.Left)
	this.doSerialize(root.Right)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.doSerialize(root)
	serializeStr := strings.Join(this.nodesVal, ",")
	fmt.Println(serializeStr)
	return serializeStr
}

// Deserializes your encoded data to tree.
func (this *Codec) doDeserialize() *TreeNode {
	head := this.nodesVal[0]
	this.nodesVal = this.nodesVal[1:]
	// poll one node
	if head == "#" {
		return nil
	}
	rootVal, _ := strconv.Atoi(head)
	root := TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	root.Left = this.doDeserialize()
	root.Right = this.doDeserialize()
	return &root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.nodesVal = strings.Split(data, ",")
	return this.doDeserialize()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	
}
