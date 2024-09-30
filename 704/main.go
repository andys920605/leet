package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx <= rightIdx {
		currentIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[currentIdx] == target {
			return currentIdx
		} else if nums[currentIdx] < target {
			leftIdx = currentIdx + 1
		} else {
			rightIdx = currentIdx - 1
		}
	}
	return -1
}
