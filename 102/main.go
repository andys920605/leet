package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 17,
			},
		},
	}
	rsp := levelOrder(root)
	fmt.Println(rsp)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := -1
	target := make(map[int][]int, 0)
	traversal(root, level, target)
	rsp := make([][]int, len(target))
	for level, v := range target {
		rsp[level] = v
	}
	return rsp
}

func traversal(node *TreeNode, level int, target map[int][]int) {
	if node == nil {
		return
	}

	level++
	if v, isExist := target[level]; !isExist {
		box := make([]int, 0)
		box = append(box, node.Val)
		target[level] = box
	} else {
		v = append(v, node.Val)
		target[level] = v
	}

	traversal(node.Left, level, target)
	traversal(node.Right, level, target)
}
