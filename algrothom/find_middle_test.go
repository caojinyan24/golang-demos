package algrothom

import (
	"fmt"
	"testing"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//示例 3：
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//示例 4：
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//示例 5：
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106

func TestFind(t *testing.T) {
	fmt.Println(Find([]int{1, 3}, []int{2}))
	fmt.Println(Find([]int{1, 2}, []int{3, 4}))
	fmt.Println(Find([]int{0, 0}, []int{0, 0}))
	fmt.Println(Find([]int{}, []int{1}))
	fmt.Println(Find([]int{2}, []int{}))
}

//判断当前的中位数在合并后的数组中的次序
//通过并序排序，组合数据，在组合到获得的次序的时候，输出
func Find(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	var index = length/2 - 1
	index1 := 0
	index2 := 0
	var nums3IndexFlag = 0
	var middle int
	for {
		if nums1[index1] > nums2[index2] {
			nums3IndexFlag = nums3IndexFlag + 1
			if nums3IndexFlag == index {
				if length%2 == 0 {
					return float64(nums1[index1])
				} else {
					middle = nums1[index1]
				}
			} else if nums3IndexFlag == index+1 {
				middle = nums1[index1] + middle
				return float64(middle / 2.0)
			}
			index1 = index1 + 1
		} else {
			nums3IndexFlag = nums3IndexFlag + 1
			if nums3IndexFlag == index {
				if length%2 == 0 {
					return float64(nums2[index2])
				} else {
					middle = nums2[index2]
				}
			} else if nums3IndexFlag == index+1 {
				middle = nums2[index2] + middle
				return float64(middle / 2.0)
			}
			index2 = index2 + 1
		}
	}
}
