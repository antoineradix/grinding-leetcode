// https://leetcode.com/problems/container-with-most-water
package main

func maxArea(height []int) int {
	max := 0
	l := len(height)

	for left, right := 0, l-1; left < right; {
		tmp := min(height[left], height[right]) * (right - left)
		if tmp > max {
			max = tmp
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}
