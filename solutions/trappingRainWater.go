package solutions

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Trap(arr []int) int {
	water := 0
	n := len(arr)

	if n > 2 {
		left := make([]int, n)
		right := make([]int, n)

		left[0] = arr[0]
		right[n-1] = arr[n-1]

		for i := 1; i < n; i++ {
			ri := n - 1 - i
			left[i] = max(arr[i], left[i-1])
			right[ri] = max(arr[ri], right[ri+1])
		}

		for i, h := range arr {
			water += min(left[i], right[i]) - h
		}
	}

	return water
}
