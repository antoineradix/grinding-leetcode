// https://leetcode.com/problems/product-of-array-except-self
package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rProduct := 1
	for i := l - 2; i >= 0; i-- {
		rProduct *= nums[i+1]
		res[i] *= rProduct
	}
	return res
}
