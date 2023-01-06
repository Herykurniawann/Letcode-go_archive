//https://leetcode.com/problems/palindrome-number/
package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var y int
	for i := x; i > 0; i /= 10 {
		y = y*10 + i%10
	}
	return x == y
}
