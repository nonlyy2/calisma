package leetcode

import (
	"math/rand"
)

// Задача №380: Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/

type RandomizedSet struct {
	nums []int
	pos  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		pos:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.pos[val]; ok {
		return false
	}
	this.pos[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.pos[val]
	if !ok {
		return false
	}

	lastIdx := len(this.nums) - 1
	lastVal := this.nums[lastIdx]

	this.nums[idx] = lastVal
	this.pos[lastVal] = idx

	this.nums = this.nums[:lastIdx]
	delete(this.pos, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.nums) == 0 {
		return 0
	}
	randIdx := rand.Intn(len(this.nums))
	return this.nums[randIdx]
}
