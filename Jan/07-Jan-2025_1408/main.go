package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"mass","as","hero","superhero"}
	fmt.Println(stringMatching(words))

}

func stringMatching(words []string) []string {
    result := []string{}

    for i, word := range words {
        for j:= 0 ; j<len(words) ; j++{
            if i != j && strings.Contains(words[j], word) {
                result = append(result, word)
                break
            }
        }
    }

    return result
}