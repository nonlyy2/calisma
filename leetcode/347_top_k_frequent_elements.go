package leetcode

// Задача №347: Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int) // stores value and freq
	for _, val := range nums {
		count[val]++ // +1 for freq of exact number
	}

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, FindMaxAndDelete(count))
	}

	return result
}

// to find key for maximum frequency and delete pair from map
func FindMaxAndDelete(count map[int]int) (maxKey int) {
	if len(count) == 0 {
		return 0
	}

	maxFreq := 0
	maxKey = 0
	for key, freq := range count {
		if freq > maxFreq {
			maxFreq = freq
			maxKey = key
		}
	}

	delete(count, maxKey)
	return maxKey
}
