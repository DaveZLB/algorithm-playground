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
	 VMAX,VS,VE = MCS(A,0,len(A)-1)
	 fmt.Printf("MCSAlgorithm,MAX=%d,(%d,%d)\n",VMAX,VS,VE)
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

 /**
 time complexity O(nlogn)
 MCS(A,s,t)
 Input: A[s...t] with s <= t
 Output: MCS of A [s...t]
 begin
 	if s = t then return A[s];
 	else
 		m <- [(s+t)/2];
 		Find MCS(A,s,m);
 		Find MCS(A,m+1,t);
 		Find MCS that contains both A[m] and A[m+1];
 		return maximum of the three sequences found
 	end
 end

 First Call: MCS(A,1,n)

  */
func MCS(A []int, s int, t int) (int, int, int) {
	VMAX := 0
	VS := -1
	VE := -1
	if s == t {
		VS = s
		VE = s
		return A[s], VS, VE
	}
	m := (s + t) / 2
	max1, VS, VE := MCS(A, s, m)
	max2, VS, VE := MCS(A, m+1, t)
	//cal max3
	lMax := 0
	lTotal := 0
	rMax := 0
	rTotal := 0
	max3 := lMax + rMax
	//cal left max val from index m to s
	for i := m; i >= s; i-- {
		lTotal += A[i]
		if lTotal > lMax {
			lMax = lTotal
			VS = i
		}
	}
	//cal right max val from index m+1 to t
	for j := m + 1; j <= t; j++ {
		rTotal += A[j]
		if rTotal > rMax {
			rMax = rTotal
			VE = j
		}
	}
	max3 = lMax + rMax
	// compare max1,max2,max3
	VMAX = max1
	if VMAX < max2 {
		VMAX = max2
	}
	if VMAX < max3 {
		VMAX = max3
	}
	return VMAX, VS, VE
}
