package main

import "fmt"

func main() {
	palindrome := "aabca"
	fmt.Println(countPalindromicSubsequence(palindrome))
	palindrome = "aabaa"
	fmt.Println(countPalindromicSubsequence(palindrome))
}


func countPalindromicSubsequence(s string) int {
    suffixCount := make(map[byte]int)
    prefixCount := make(map[byte]bool)
    visited := make(map[byte]map[byte]bool)

    for i := 0; i < len(s); i++ {
        suffixCount[s[i]]++
    }

    count := 0

    for i := 0; i < len(s); i++ {
        currentChar := s[i]

        suffixCount[currentChar]--
        if suffixCount[currentChar] == 0 {
            delete(suffixCount, currentChar)
        }

        if _, exists := visited[currentChar]; !exists {
            visited[currentChar] = make(map[byte]bool)
        }

        for x := byte('a'); x <= byte('z'); x++ {
            if prefixCount[x] && suffixCount[x] > 0 && !visited[currentChar][x] {
                count++
                visited[currentChar][x] = true
            }
        }

        prefixCount[currentChar] = true
    }

    return count
}
