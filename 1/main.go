package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 7}
	target := 13

	fmt.Println(twoSum(nums, target))

}

// O(n^2)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		want := target - v
		if idx, ok := m[want]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return []int{}
}
