package main

import (
	"fmt"
	"strings"
	
)

/*	
 * 1. Two Sum
 
 Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
 You may assume that each input would have exactly one solution, and you may not use the same element twice.
 You can return the answer in any order.
  
 Example 1:
 Input: nums = [2,7,11,15], target = 9
 Output: [0,1]
 Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

 Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
 */

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
