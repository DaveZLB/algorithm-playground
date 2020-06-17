package main

import "fmt"

func main(){
	A := []int{9,5,1,7,8,2}
	result := insertSort(A)
	for _,v :=range result{
		fmt.Printf("%d,",v)
	}

}

/**
Input: A[1...n] is an array of numbers
for j <- 2 to n do
	key <- A[j];
	i <- j - 1;
	while i >= 1 And A[i] > key do
		A[i+1] <- A[i];
		i <- i-1;
	end
	A[i+1] <- key;
	<--sorted-->key<---unsorted--->
end
 */
func insertSort(array []int) []int {
	for j := 1; j < len(array); j++ {
		i := j - 1
		temp := array[j]
		for i >= 0 && array[i] > temp {
			array[i+1] = array[i]
			i = i - 1
		}
		array[i+1] = temp
	}
	return array
}