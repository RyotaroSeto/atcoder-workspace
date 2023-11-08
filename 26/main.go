package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

// func removeDuplicates(nums []int) int {
// 	i := 0
// 	for j := range nums {
// 		if nums[i] != nums[j] {
// 			i += 1
// 			nums[i] = nums[j]
// 		}
// 	}
// 	return i + 1
// }

// func removeDuplicates(nums []int) int {
// 	prev := nums[0]
// 	l := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != prev {
// 			nums[l] = nums[i]
// 			l++
// 		}
// 		prev = nums[i]
// 	}
// 	return l
// }
