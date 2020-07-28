package main

import (
	"fmt"
)

func main(){
	A := []int{14,7,18,3,10,19,11,23,2,25,16,17}
	r := BruteForceAlgorithm(A)
	fmt.Println(r)
	r1 := SortAndCount(A,0,len(A)-1)
	println(r1)
}

func BruteForceAlgorithm(A []int) int {
	r := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				r += 1
			}
		}
	}
	return r
}

/**
Sort-and-Count(L)
Input: L
Output: r_l,L
if L is empty then
	return 0,L;
end
Divide L into two halves A and B;
(r_A,A) <- Sort-and-Count(A);
(r_B,B) <- Sort-and-Count(B);
(r_L,L) <- Merge-and-Sort(A,B);
return r_A + r_B + r_L, L;
 */
func SortAndCount(A []int,left int, right int) int{
	if left < right {
		center := (left + right) / 2
		l_r := SortAndCount(A,left,center)
		r_r := SortAndCount(A,center+1,right)
		m_r := 0
		//merge
		leftLength := center - left + 1
		rightLength := right-center
		leftArray := make([]int,leftLength)
		rightArray := make([]int, rightLength)
		for i := 0;i < leftLength; i++{
			leftArray[i] = A[i+left]
		}
		for j :=0 ; j < rightLength; j++{
			rightArray[j] = A[j+center+1]
		}
		i := 0
		j := 0
		k := 0
		for i < leftLength && j < rightLength{
			if leftArray[i] < rightArray[j]{
				A[left + k] = leftArray[i]
				i++
			}else{
				A[left + k] = rightArray[j]
				j++
				m_r += leftLength - i
			}
			k++
		}
		for i < leftLength {
			A[left + k] = leftArray[i]
			i++
			k++
		}
		for j < rightLength {
			A[left+k] = rightArray[j]
			j++
			k++
		}
		return l_r + m_r + r_r
	}
	return 0
}
