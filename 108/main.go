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
	nums := []int{-10, -3, 0, 5, 9}
	rsp := sortedArrayToBST(nums)
	j1, _ := json.Marshal(rsp)
	fmt.Println(string(j1))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := len(nums) / 2
	value := nums[index]
	root := &TreeNode{
		Val: value,
	}
	if len(nums) == 1 {
		return root
	}
	left := nums[0:index]
	right := nums[index+1:]
	root.Left = sortedArrayToBST(left)
	root.Right = sortedArrayToBST(right)

	return root
}
