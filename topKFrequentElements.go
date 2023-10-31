// https://leetcode.com/problems/top-k-frequent-elements
package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num] += 1
	}

	type Pair struct {
		value int
		freq  int
	}

	pairs := make([]Pair, 0, len(countMap))

	for value, freq := range countMap {
		pairs = append(pairs, Pair{value, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = pairs[i].value
	}
	return ret
}
