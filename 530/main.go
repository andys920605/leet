package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// root := &TreeNode{
	// 	Val: 4,
	// 	Right: &TreeNode{
	// 		Val: 6,
	// 	},
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Right: &TreeNode{
	// 			Val: 3,
	// 		},
	// 		Left: &TreeNode{
	// 			Val: 1,
	// 		},
	// 	},
	// }
	root2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	rsp := getMinimumDifference(root2)
	fmt.Println(rsp)
}

var box []int

func getMinimumDifference(root *TreeNode) int {
	box = []int{}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	traverse(root)
	return box[1] - box[0]
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left)
	box = append(box, node.Val)
	if len(box) == 2 {
		return
	}
	traverse(node.Right)
	if len(box) == 2 {
		return
	}
}
