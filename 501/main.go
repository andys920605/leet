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
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	rsp := findMode(root)

	j1, _ := json.Marshal(rsp)
	fmt.Println(string(j1))
}

func findMode(root *TreeNode) []int {
	collection := make(map[int]int)
	traversal(root, collection)
	target := []int{}
	num := 0
	value := 0
	for i, v := range collection {
		if v > num {
			num = v
			value = i
		}
	}
	target = append(target, value)
	return target
}

func traversal(root *TreeNode, collection map[int]int) {
	if root == nil {
		return
	}
	collection[root.Val]++
	traversal(root.Left, collection)
	traversal(root.Right, collection)
}
