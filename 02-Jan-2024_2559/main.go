package main

import (
	"fmt"
)

func main() {
	words := []string{"aba", "baba", "aba", "xzxb"}
	queries := [][]int{{1, 1}, {2, 2}, {3, 3}, {1, 3}}
	fmt.Println(vowelStrings(words, queries))
}

func vowelStrings(words []string, queries [][]int) []int {
    vowels := map[byte]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        }

    prefix_count := make([]int, len(words)+1)
    prev := 0

    for i, w := range words {
        if vowels[w[0]] && vowels[w[len(w)-1]] {
            prev ++
        }
        prefix_count[i+1] = prev
    }

    result := make([]int, len(queries))

    for i, q := range queries {
        result[i] = prefix_count[q[1] + 1] - prefix_count[q[0]]
    }

    return result
}