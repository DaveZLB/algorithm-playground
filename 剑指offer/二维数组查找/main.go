package main

import "fmt"

//题目描述：
//给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。
// 给定一个数，判断这个数是否在该二维数组中。
//要求时间复杂度 O(M + N)，空间复杂度 O(1)。其中 M 为行数，N 为 列数。

//Consider the following matrix:
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//
//Given target = 5, return true.
//Given target = 20, return false.

//解题思路：
//该二维数组中的一个数，小于它的数一定在其左边，大于它的数一定在其下边。
// 因此，从右上角开始查找，就可以根据 target 和当前元素的大小关系来缩小查找区间，当前元素的查找区间为左下角的所有元素

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	existed := find(matrix, 5)
	fmt.Printf("5 -> %v\n", existed)
	existed = find(matrix, 20)
	fmt.Printf("20 -> %v", existed)
}

func find(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	r := 0
	c := cols - 1
	for r <= rows-1 && c >= 0 {
		if target == matrix[r][c] {
			return true
		} else if target < matrix[r][c] {
			c--
		} else if target > matrix[r][c] {
			r++
		}
	}
	return false
}
