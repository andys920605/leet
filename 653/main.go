package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	k := 9
	rsp := findTarget(root, k)

	j1, _ := json.Marshal(rsp)
	fmt.Println(string(j1))
}

func findTarget(root *TreeNode, k int) bool {
	box := make(map[int]struct{})
	return traverse(root, k, box)
}

func traverse(node *TreeNode, target int, box map[int]struct{}) bool {
	if node == nil {
		return false
	}
	if _, isExist := box[node.Val]; isExist {
		return true
	}
	want := target - node.Val
	box[want] = struct{}{}
	if status := traverse(node.Left, target, box); status {
		return true
	}
	if status := traverse(node.Right, target, box); status {
		return true
	}
	return false
}
