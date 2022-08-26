# golang_sliding_window_maximum

You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.

## Examples

**Example 1:**

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7      3
 1 [3  -1  -3] 5  3  6  7      3
 1  3 [-1  -3  5] 3  6  7      5
 1  3  -1 [-3  5  3] 6  7      5
 1  3  -1  -3 [5  3  6] 7      6
 1  3  -1  -3  5 [3  6  7]     7
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]

```

**Constraints:**

- `1 <= nums.length <= $10^6$`
- $`-10^4$ <= nums[i] <= $10^4$`
- `1 <= k <= nums.length`

## 解析

題目給定一個整數陣列 nums ， 還有一個正整數 k 代表 slide window 的大小

要求寫一個演算法 算出 nums 把 slide-window 由 nums 最左到最右界走過一遍 找出每個位置 slide-window 內的最大值

首先要知道找從左走到右會有多少個

假設 n = len(nums),  slide-window 大小是 k

則最後一個 slide window 起點會在nums 的 n - k  位置

如下圖：

![](https://i.imgur.com/uuSuV42.png)

而 sliding-window 的第一個起點在 0

所以從 0 到 n-k 公有 n-k + 1 個 slide-window maximum 值

要怎麼透過 sliding-window 特性每次找到在 window 內的最大值呢？

我們可以設計一個 queue 每次依序紀錄  相鄰位置的索引值在 queue 

且queue 只能放 k 個鄰近的值

具體來說 當 i 是當下位置 且 i ≥ k， 則  queue 的最左界不能是 i - k

當超過時，需要把最左界往右移 

![](https://i.imgur.com/PrnQL3U.png)

要找到最大值代表需要在queue 內需要把目前小 nums[i] 的值移出

這樣就可以一直維持在後面放入的值比前面大

![](https://i.imgur.com/Sl8o0is.png)

也就是可以維持每次 queue內的第一個都是 window 內的最大值

所以作法如下

每次當 i 走到超過 slide windows size 大小時 檢查 window 左界有沒有小於 i - k

如果是則 把左界移出 window

當 window 內有值時，把 小於 nums[i] 的 index 移出 window

把 nums[i] 放入 window

當 i 走到大於等於 k-1 位置時

把 windows[0] 放入 result

具體作法如下：

![](https://i.imgur.com/2flGsui.png)

## 程式碼
```go
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

```
## 困難點

1. 要知道如何透過 slide-window 特性找到區域最大
2. 需要知道 slide window 起始點

## Solve Point

- [x]  初始化 window = []int{}, result := []int{}
- [x]  對所有 i := 0..len(nums)-1 做以下運算
- [x]  當 i ≥ k && window[0] ≤ i - k 時，代表左界超過範圍需要做以下運算
- [x]  更新 window = window[1: len(window)] // 把左界移出 window
- [x]  把所有小於 nums[i] 的索引值都移出window, 當 len(window) > 0 && nums[window[len(window)-1]] ≤ nums[i], 更新 window = windows[:len(window)-1]
- [x]  把目前 index 加入 window , 更新 window = append(window, i)
- [x]  當 i ≥ k -1 時，代表第1個 slide-window 的值都蒐集完了做以下運算
- [x]  result = append(result, num[windows[0]]
- [x]  最後回傳 result