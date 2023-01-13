//https://leetcode.com/problems/roman-to-integer/
package main

func romanToInt(s string) int {
	var r int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i < len(s)-1 && (s[i+1] == 'V' || s[i+1] == 'X') {
				r -= 1
			} else {
				r += 1
			}
		case 'V':
			r += 5
		case 'X':
			if i < len(s)-1 && (s[i+1] == 'L' || s[i+1] == 'C') {
				r -= 10
			} else {
				r += 10
			}
		case 'L':
			r += 50
		case 'C':
			if i < len(s)-1 && (s[i+1] == 'D' || s[i+1] == 'M') {
				r -= 100
			} else {
				r += 100
			}
		case 'D':
			r += 500
		case 'M':
			r += 1000
		}
	}
	return r
}
