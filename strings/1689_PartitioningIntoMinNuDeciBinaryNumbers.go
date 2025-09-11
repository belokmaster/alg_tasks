package main

func minPartitions(n string) int {
	nums := 0
	for i := 0; i < len(n); i++ {
		curr := int(n[i] - '0')
		if nums < curr {
			nums = curr
		}
	}
	return nums
}
