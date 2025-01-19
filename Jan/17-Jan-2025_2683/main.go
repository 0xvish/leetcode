package main

import "fmt"

func main() {
    fmt.Println(doesValidArrayExist([]int{2,3,1,6,7}));
}

func doesValidArrayExist(derived []int) bool {
    arr_xor := 0;
    for  i :=0; i<len(derived); i++ {
        arr_xor ^= derived[i];
    }
    return arr_xor==0;
}