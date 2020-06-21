package main

import "fmt"

//题目：
//在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。
// 请找出数组中任意一个重复的数字。

// 要求：
//时间复杂度 O(N)，空间复杂度 O(1)。因此不能使用排序的方法，也不能使用额外的标记数组。

//思路：
//对于这种数组元素在 [0, n-1] 范围内的问题，
// 可以将值为 i 的元素调整到第 i 个位置上进行求解。
// 本题要求找出重复的数字，因此在调整过程中，如果第 i 位置上已经有一个值为 i 的元素，就可以知道 i 值重复。

func main() {
	//以 (2, 3, 1, 0, 2, 5) 为例
	origin := []int{2, 3, 1, 0, 2, 5}
	r := make([]int, 1)
	existed := deduplicate(origin, r)
	if existed {
		fmt.Printf("Duplicate element is %d", r[0])
	} else {
		fmt.Println("Duplicate element not existed")
	}
}

func deduplicate(origin []int, result []int) bool {
	for i := 0; i < len(origin); i++ {
		for origin[i] != i {
			if origin[i] == origin[origin[i]] {
				result[0] = origin[i]
				return true
			}
			//交换位置i和origin[i]位置的元素
			origin[i], origin[origin[i]] = origin[origin[i]], origin[i]
		}
	}
	return false
}
