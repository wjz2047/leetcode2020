package TwoSum

import "sort"

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
