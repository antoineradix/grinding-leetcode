// https://leetcode.com/problems/longest-consecutive-sequence
package main

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for _, n := range nums {
		if numSet[n] == false {
			continue
		}
		tmpStreak := 0
		for i := n - 1; numSet[i]; i-- {
			numSet[i] = false
			tmpStreak++
		}
		for i := n; numSet[i]; i++ {
			numSet[i] = false
			tmpStreak++
		}

		if tmpStreak > longestStreak {
			longestStreak = tmpStreak
		}
	}
	return longestStreak
}
