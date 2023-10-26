// https://leetcode.com/problems/valid-palindrome
package main

func isAlphaNum(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func cmp(a byte, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	if a == b {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left <= right; {
		if isAlphaNum(s[left]) == false {
			left++
		} else if isAlphaNum(s[right]) == false {
			right--
		} else {
			if cmp(s[left], s[right]) == false {
				return false
			}
			left++
			right--
		}
	}
	return true
}
