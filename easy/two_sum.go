package main

import (
	"fmt"
)

// Calculate the two sum of a given array of integers and a target value.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)  // Initialize a map to store seen numbers and their indices
	for i, num := range nums {
		complement := target - num // Calculate the complement of the current number
		
		if idx, found := seen[complement]; found {
					return []int{idx, i}
				}
				seen[num] = i
	}
	return nil
}

// main function to test the twoSum function
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
