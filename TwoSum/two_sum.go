package TwoSum

import (
	"math"
	"sort"
)

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v, ok := mp[target - nums[i]]
		if ok {
			return []int{v,i}
		}
		mp[nums[i]] = i
	}
	return []int{}
}
// numbers sorted
func TwoSum2(numbers []int, target int) []int {
	lth := len(numbers)
	low := 0
	high := lth -1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low+1, high+1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{}
}


func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	lth := len(nums)
	sort.Ints(nums)
	for i := 0; i <= lth - 3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			j := i + 1
			k := lth -1
			for j < k {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					for j+ 1 < lth && nums[j+1] == nums[j] {
						j++
					}
					j++
					for k-1 > 0 && nums[k-1] == nums[k] {
						k--
					}
					k--
				} else if sum < 0 {
					j++
				} else {
					k--
				}
			}
		}
	}
	return result
}

func ThreeSumClosest(nums []int, target int) int {
	min := math.MaxInt64
	var result int
	lth := len(nums)
	sort.Ints(nums)
	for i := 0; i <= lth-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			j := i+1
			k := lth-1
			for j < k {
				sum := nums[i] + nums[j] + nums[k]
				diff := sum - target
				if diff == 0 {
					return sum
				} else if diff > 0 {
					k--
					if diff < min {
						min = diff
						result = sum
					}
				} else {
					j++
					if -diff < min {
						min = -diff
						result = sum
					}
				}
			}
		}
	}
	return result
}
