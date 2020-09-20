package TwoSum

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
