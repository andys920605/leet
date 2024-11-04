package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	box := make([]int, len(nums))
	for i, v := range nums {
		box[i] = v * v
	}
	sort.Slice(box, func(i, j int) bool {
		return box[i] < box[j]
	})
	return box
}
