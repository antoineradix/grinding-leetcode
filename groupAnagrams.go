// https://leetcode.com/problems/group-anagrams
package main

func groupAnagrams(strs []string) [][]string {
	h := make(map[[26]byte][]string)

	for _, str := range strs {
		key := [26]byte{}
		for _, c := range []rune(str) {
			key[c-'a'] += 1
		}
		h[key] = append(h[key], str)
	}

	ret := [][]string{}
	for _, anagrams := range h {
		ret = append(ret, anagrams)
	}
	return ret
}
