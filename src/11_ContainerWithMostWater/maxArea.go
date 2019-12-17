package _1_ContainerWithMostWater


func maxArea(height []int) int {
	var (
		p, q int
		max, area int
	)
	if len(height) <= 1 {
		return 0
	}
	max = 0
	p = 0
	q = len(height) - 1
	for q > p {
		if height[q] > height[p] {
			area = height[p] * (q - p)
			p++
		} else {
			area = height[q] * (q - p)
			q--
		}
		if area > max {
			max = area
		}
	}
	return max
}

func MaxArea(height []int) int {
	return maxArea(height)
}