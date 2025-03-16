package main

import "math"

func repairCars(ranks []int, cars int) int64 {
	left := 0
	right := ranks[0] * cars * cars

	for left < right {
		mid := left + (right-left)/2
		total := 0

		for _, r := range ranks {
			// Для каждого механика вычисляется, сколько автомобилей он может починить за время
			total += int(math.Sqrt(float64(mid / r)))
			if total >= cars {
				break
			}
		}

		if total >= cars {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}
