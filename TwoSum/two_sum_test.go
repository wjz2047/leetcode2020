package TwoSum

import (
	"testing"
	"fmt"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3,3}
	target := 6
	fmt.Println(TwoSum(nums, target))
}

func TestTwoSum2(t *testing.T) {
	nums := []int{-1,0}
	target := -1
	fmt.Println(TwoSum2(nums, target))
}
