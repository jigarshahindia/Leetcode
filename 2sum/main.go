package main

import "fmt"

/*
https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

func sum(num []int, target int) []int {
	for i, x := range num {
		for j, y := range num {
			if i == j {
				continue
			}
			if x+y == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	target := 11
	num := []int{1, 4, 5, 10}
	fmt.Println(sum(num, target))
}
