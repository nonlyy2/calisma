package leetcode

// Задача №33: Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	if len(nums) == 1 && nums[0] != target {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // левая часть норм
			if nums[left] <= target && target < nums[mid] { // таргет в левой части
				right = mid - 1
			} else { // таргет в правой части
				left = mid + 1
			}
		} else { // правая часть норм
			if nums[mid] < target && target <= nums[right] { // таргет в правой части
				left = mid + 1
			} else { // таргет в левой части
				right = mid - 1
			}
		}
	}

	return -1
}
