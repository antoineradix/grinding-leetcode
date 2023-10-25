// https://leetcode.com/problems/contains-duplicate
package main

func containsDuplicate(nums []int) bool {
	checkDup := make(map[int]bool)
	for _, v := range nums {
		if checkDup[v] {
			return true
		}
		checkDup[v] = true
	}
	return false
}
