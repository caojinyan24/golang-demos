package algrothom

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
//
//
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//示例 2：
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//示例 3：
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//提示：
//
//2 <= nums.length <= 104
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//只会存在一个有效答案
//进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
//////////////////////////////////////////////////////////////////////////////////////////////////////////////
//思路：每次取出来一个，在剩余的元素中找是否有匹配的，如果没有，则从这个数组中去掉这个值
//输出的时候需要使用序号，做一个map，获取到值的时候从里边取对应的序号输出（存在重复的，则保存一个数组）
func findSumTarget(array []int, target int) []int {
	// make index map
	var indexMap = make(map[int][]int)
	for index, item := range array {
		value, ok := indexMap[item]
		if ok {
			indexMap[item] = append(value, index)
		} else {
			indexMap[item] = []int{index}
		}
	}
	//get value
	for index, item := range array {
		if target-item == item {
			if len(indexMap[item]) <= 1 {
				continue
			} else {
				targetIndexArray := indexMap[item]
				return targetIndexArray[0:2]
			}
		}
		deductedValue, ok := indexMap[target-item]
		if ok && len(deductedValue) > 0 {
			return []int{index, deductedValue[0]}
		}
	}
	return []int{}
}
func TestFindSumTarget(t *testing.T) {
	fmt.Println(findSumTarget([]int{2, 7, 11, 15}, 9))
	fmt.Println(findSumTarget([]int{3, 2, 4}, 6))
	fmt.Println(findSumTarget([]int{3, 3}, 6))
}
