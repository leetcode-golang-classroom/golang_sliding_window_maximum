package sol

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || n < k {
		return make([]int, 0)
	}
	result, window := []int{}, []int{}
	for i, v := range nums {
		if i >= k && window[0] <= i-k {
			// shift window
			window = window[1:]
		}
		// remove all smaller nums[i] index in window
		for len(window) > 0 && nums[window[len(window)-1]] < v {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}
