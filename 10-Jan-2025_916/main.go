package main

import (
	"fmt"
	"strings"
)
func main() {
    words1 := []string{"amazon","apple","facebook","google","leetcode"}
    words2 := []string{"e","o"}
    fmt.Println(wordSubsets(words1, words2))
}


func wordSubsets(words1 []string, words2 []string) []string {
    answer := []string{}
    count_2 := make(map[byte]int)

    for _, word := range words2 {
        tempCount := make(map[byte]int)
        for i := 0; i < len(word); i++ {
            tempCount[word[i]]++
            if tempCount[word[i]] > count_2[word[i]] {
                count_2[word[i]] = tempCount[word[i]]
            }
        }
    }

    for _, word := range words1 {
        isValid := true
        for char, requiredCount := range count_2 {
            if strings.Count(word, string(char)) < requiredCount {
                isValid = false
                break
            }
        }
        if isValid {
                answer = append(answer, word)
        }
    }

    return answer
}