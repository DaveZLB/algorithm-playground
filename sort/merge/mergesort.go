package main

import "fmt"

func main() {
	A := []int{9, 5, 1, 7, 8, 2}
	mergeSort(A, 0, len(A)-1)
	for _, v := range A {
		fmt.Printf("%d,", v)
	}

}

/**
mergeSort(A,left,right)
if left < right then
	center <- (left + right)/2
	mergeSort(A,left,center)
	mergeSort(A,center+1,right)
	Merge two sorted arrays
end
 */

func mergeSort(array []int, left int, right int) {
	if left < right {
		center := (left + right) / 2
		mergeSort(array, left, center)
		mergeSort(array, center+1, right)
		merge(array, left, right, center)
	}
}

func merge(array []int, left int, right int, center int) {
	//copy left and right array to new array
	leftLen := center - left + 1
	rightLen := right - center
	leftArray := make([]int, leftLen)
	rightArray := make([]int, rightLen)
	for i := 0; i < leftLen; i++ {
		leftArray[i] = array[left+i]
	}
	for j := 0; j < rightLen; j++ {
		rightArray[j] = array[center+1+j]
	}
	// merge
	i := 0
	j := 0
	k := 0
	//When both leftArray and rightArray are not traversed
	for i < leftLen && j < rightLen {
		if leftArray[i] < rightArray[j] {
			array[left+k] = leftArray[i]
			i++
		} else {
			array[left+k] = rightArray[j]
			j++
		}
		k++
	}
	//When leftArray not traversed
	for i < leftLen {
		array[left+k] = leftArray[i]
		i++
		k++
	}
	//When rightArray not traversed
	for j < rightLen {
		array[left+k] = rightArray[j]
		j++
		k++
	}

}
