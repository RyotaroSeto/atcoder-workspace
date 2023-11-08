package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
