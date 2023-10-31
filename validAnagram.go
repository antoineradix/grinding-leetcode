// https://leetcode.com/problems/valid-anagram
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var key_a [26]int
	var key_b [26]int
	for _, c := range []rune(s) {
		key_a[c-'a'] += 1
	}
	for _, c := range []rune(t) {
		key_b[c-'a'] += 1
	}
	if key_a != key_b {
		return false
	}
	return true
}
