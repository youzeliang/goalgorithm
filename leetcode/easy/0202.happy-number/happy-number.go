package _202_happy_number

//# [202. Happy Number](https://leetcode.com/problems/happy-number/)
//
//## 题目
//Write an algorithm to determine if a number is "happy".
//
//A happy number is a number defined by the following process:
//1. Starting with any positive integer, replace the number by the sum of the squares of its digits,
//1. and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
//
//Example: 19 is a happy number
//```
//1^2 + 9^2 = 82
//8^2 + 2^2 = 68
//6^2 + 8^2 = 100
//1^2 + 0^2 + 0^2 = 1
//```
//
//Credits:Special thanks to @mithmatt and @ts for adding this problem and creating all test cases.
//

func isHappy(n int) bool {
	// 开一个 map 判断是否循环
	hash := make(map[int]bool)
	for n!=1 {
		// 出现了循环，证明不是快乐数
		if _,ok := hash[n]; ok {
			return false
		}
		hash[n] = true
		next := 0
		// 计算下一个数字
		for n!=0 {
			next += (n%10) * (n%10)
			n /= 10
		}
		n = next
	}
	return true
}