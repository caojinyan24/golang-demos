package algrothom

import (
	"fmt"
	"testing"
)

//3. Longest Substring Without Repeating Characters
//Medium
//
//19225
//
//879
//
//Add to List
//
//Share
//Given a string s, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//Example 4:
//
//Input: s = ""
//Output: 0
//
//Constraints:
//
//0 <= s.length <= 5 * 104
//s consists of English letters, digits, symbols and spaces.
func TestFindLongestStr(t *testing.T) {
	fmt.Println(FindLongestStr([]rune{'a', 'b', 'c', 'a', 'b', 'c', 'b', 'b'}))
	fmt.Println(FindLongestStr([]rune{'b', 'b', 'b', 'b', 'b'}))
	fmt.Println(FindLongestStr([]rune{'p', 'w', 'w', 'k', 'e', 'w'}))
	fmt.Println(FindLongestStr([]rune{}))

}
func FindLongestStr(str []rune) int {
	if len(str) <= 1 {
		return len(str)
	}
	var itemMap = make(map[rune]bool)
	for _, item := range str {
		itemMap[item] = true
	}
	for offset := len(str) - len(itemMap); offset <= len(str); offset = offset + 1 {
		for startFlag := 0; -offset+startFlag <= 0; startFlag = startFlag + 1 {
			if !checkIfExistSame(str[startFlag : len(str)-offset+startFlag]) {
				return len(str) - offset
			}
		}
	}
	return 1
}
func checkIfExistSame(str []rune) bool {
	var valueMap = make(map[rune]bool)
	for _, item := range str {
		if valueMap[item] {
			return true
		} else {
			valueMap[item] = true
		}
	}
	return false
}
