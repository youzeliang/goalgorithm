package sort

// 寻找全排列的下一个数
// 给一个正整数，找出正整数的所有数字全排列的下一个数

// 解题步骤
// 1. 从后向前查看逆序区，找到逆序区的前一位，也就是数字置换的边界
// 2. 让逆序区域的前一位和逆序区的中的大于它的最小数字交换位置
// 3. 把原来的逆序区转化为顺序状态
