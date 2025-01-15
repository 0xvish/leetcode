package main

import (
	"fmt"
)

func main() {
	A := []int{2, 3, 4, 5, 6}
	B := []int{1, 2, 3, 4, 5}
	fmt.Println(findThePrefixCommonArray(A, B))
}

func findThePrefixCommonArray(A []int, B []int) []int {
    C := make([]int, len(A))

    for i:=0 ; i<len(A) ; i++ {
        found := false
        for j:=0 ; j<len(B) ; j++ {
            if A[i] == B[j] {
                found = true
            }
            if found && j>=i {
                C[j]++
            }

        }
    }

    return C
}