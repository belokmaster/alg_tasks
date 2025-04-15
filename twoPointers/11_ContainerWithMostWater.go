package main

func maxArea(height []int) int {
	ans := 0
	n := len(height)
	left := 0
	right := n - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
