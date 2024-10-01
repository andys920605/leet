package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	count := 0
	for _, num := range nums {
		if num != val {
			nums[count] = num
			count++
		}
	}
	return count
}
