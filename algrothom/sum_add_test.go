package algrothom

import (
	"fmt"
	"testing"
)

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
//示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//示例 2：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//示例 3：
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//提示：
//
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零

func TestSumAdd(t *testing.T) {
	fmt.Println(sumAdd([]int{2, 4, 3}, []int{5, 6, 4}))
	fmt.Println(sumAdd([]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}))
	fmt.Println(sumAdd([]int{0}, []int{0}))

	fmt.Println(sumAddByArray([]int{2, 4, 3}, []int{5, 6, 4}))
	fmt.Println(sumAddByArray([]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}))
	fmt.Println(sumAddByArray([]int{0}, []int{0}))

}

//
func sumAddByArray(arr1 []int, arr2 []int) []int {
	var resultArray = make([]int, 0)
	var index = 0
	for {
		var sumed = 0
		if len(arr1) > index {
			sumed = sumed + arr1[index]
		}
		if len(arr2) > index {
			sumed = sumed + arr2[index]
		}
		if len(resultArray) > index {
			sumed = sumed + resultArray[index]
		} else {
			resultArray = append(resultArray, 0) //提前把当前位置放入数据
		}

		//add data to resultArray
		resultArray[index] = sumed % 10
		if sumed >= 10 {
			resultArray = append(resultArray, sumed/10)
		}
		index = index + 1
		if len(arr1)-1 < index && len(arr2)-1 < index && len(resultArray)-1 < index {
			break
		}
	}
	return resultArray
}
func sumAdd(arr1 []int, arr2 []int) []int {
	var num1 = 0
	var num2 = 0
	var poFlag = 1
	for index, item := range arr1 {
		if index == 0 {
			poFlag = 1
		} else {
			poFlag = poFlag * 10
		}
		num1 = item*poFlag + num1
	}
	for index, item := range arr2 {
		if index == 0 {
			poFlag = 1
		} else {
			poFlag = poFlag * 10
		}
		num2 = item*poFlag + num2
	}
	var sum = num2 + num1
	var output []int
	poFlag = 10
	for {
		var tempItem = sum % poFlag
		output = append(output, tempItem)
		sum = (sum - tempItem) / 10
		if sum == 0 {
			break
		}
	}
	return output
}
