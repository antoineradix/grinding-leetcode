// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
package main

func twoSum(numbers []int, target int) []int {
	l := len(numbers)

	for i := 0; i < l; i++ {
		left := i + 1
		right := l - 1

		for left <= right {
			mid := (left + right) >> 1
			sum := numbers[i] + numbers[mid]
			if sum == target {
				return []int{i + 1, mid + 1}
			}
			if sum < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nil
}
