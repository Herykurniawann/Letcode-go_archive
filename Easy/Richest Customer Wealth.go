//https://leetcode.com/problems/richest-customer-wealth/
package main

func maximumWealth(accounts [][]int) int {
	var r int
	for _, v := range accounts {
		var s int
		for _, v := range v {
			s += v
		}
		if s > r {
			r = s
		}
	}
	return r

}
