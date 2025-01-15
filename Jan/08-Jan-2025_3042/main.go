package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"a","aba","ababa","aa"}
	fmt.Println(countPrefixSuffixPairs(words))
}

func countPrefixSuffixPairs(words []string) int {
    answer := 0

    for i, word := range words {
        for j:=i+1 ; j < len(words) ; j++ {
            if i != j && isPrefixAndSuffix(word, words[j]){
                answer ++
            }
        }
    }

    return answer
}

func isPrefixAndSuffix(s1 string, s2 string) bool {
    return strings.HasPrefix(s2, s1) && strings.HasSuffix(s2, s1)
}