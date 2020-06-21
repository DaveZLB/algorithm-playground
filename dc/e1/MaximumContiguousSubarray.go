package main

import "fmt"

/**
	PROFIT HISTORY
Year	1	2	3	4	5	6	7	8	9
Profit	-3	2	1	-4	5	2	-1	3	-1

Between which two years profit is maximized ?
 */



 func main(){
 	A := []int{-3,2,1,-4,5,2,-1,3,-1}
 	VMAX,VS,VE := BruteForceAlgorithm(A)
 	fmt.Printf("BruteForceAlgorithm,MAX=%d,(%d,%d)\n",VMAX,VS,VE)
	 VMAX,VS,VE = DataReuseAlgorithm(A)
	 fmt.Printf("DataReuseAlgorithm,MAX=%d,(%d,%d)\n",VMAX,VS,VE)
 }

 /**
 time complexity: O(n^3)
 VMAX <- A[1];
 for i <- 1 to n do
 	for j <- i to n do
 		V<- 0;
 		for k <- i to j do
 			V <- V + A[k];
 		end
 		if V > VMAX then
 			VMAX <- V;
 		end
 	end
 end
 return VMAX
  */
 func BruteForceAlgorithm(A []int) (int,int,int){
 	VMAX := 0
 	VS := -1
 	VE := -1
 	for i := 0; i < len(A); i++ {
 		for j := i; j < len(A); j++{
 			V := 0
 			for k := i; k <= j ; k++{
 				V = V + A[k]
			}
 			if V > VMAX{
 				VMAX = V
 				VS = i
 				VE = j
			}
		}
	}
 	return VMAX,VS,VE
 }

 /**
 time complexity O(n^2)
 VMAX=A[1]
 for i <- 1 to n do
 	V = 0
 	for j <- i to n do
 		V <- V + A[j]
 		if V > VMAX then
 			VMAX <- V
 		end
 	end
 end
  */
 func DataReuseAlgorithm(A []int)(int,int,int){
	 VMAX := 0
	 VS := -1
	 VE := -1
	 for i := 0; i < len(A); i++ {
		 V := 0
		 for j := i; j < len(A); j++{
			 V = V + A[j]
			 if V > VMAX{
				 VMAX = V
				 VS = i
				 VE = j
			 }
		 }
	 }
	 return VMAX,VS,VE
 }