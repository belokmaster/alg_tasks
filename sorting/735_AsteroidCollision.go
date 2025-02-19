package main

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {
		alive := true

		for alive && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]

			if top < -asteroid {
				stack = stack[:len(stack)-1]
				continue
			} else if top == -asteroid {
				stack = stack[:len(stack)-1]
			}

			alive = false
		}

		if alive {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
