package main

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 4410 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	var (
		leftMax  = make([]int, len(height))
		rightMax = make([]int, len(height))
		total    = 0
	)

	for i := 1; i < len(height)-1; i++ {
		leftMax[i] = Max(leftMax[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], height[i+1])
	}

	for i := 1; i < len(height)-1; i++ {
		min := Min(leftMax[i], rightMax[i])
		if min > height[i] {
			total += min - height[i]
		}

	}
	return total
}

//同报目录下存在
//func Max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
