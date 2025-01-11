package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("annabelle", 2))
}

func canConstruct(s string, k int) bool {
    if len(s) < k {
        return false
    }
    charCount := make(map[byte]int)
    oddCount := 0

    for i:=0 ; i<len(s) ; i++ {
        charCount[s[i]]++
    }

    for _, count := range charCount {
        if (count%2) != 0 {
            oddCount++
        }
    }

    if oddCount > k {
        return false
    }

    return true
}