package problem0011

func maxArea(height []int) int {
	mx := 0
	i, j := 0, len(height) - 1
	var vol int
	for i < j {
		if height[i] > height[j] {
			vol = (j - i) * height[j]
			j--
		} else {
			vol = (j - i) * height[i]
			i++
		}
		if mx < vol {
			mx = vol
		}
	}
	return mx
}