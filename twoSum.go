// https://leetcode.com/problems/two-sum
package main

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		idx, exists := indexMap[(target - num)]
		if exists {
			return []int{idx, i}
		}
		indexMap[num] = i
	}
	return []int{0}
}
