package main

import "fmt"

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

*/
func lengthOfLongestSubstring(s string) int {
	max := 0
	for pos, char := range s {
		m := make(map[int32]bool)
		m[char] = true
		count := 1
		for inPos, inChar := range s {
			if inPos <= pos {
				continue
			}
			_, ok := m[inChar]
			if ok {
				break
			} else {
				count++
				m[inChar] = true
			}
		}
		if count > max {
			max = count
		}

	}
	return max
}

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}
