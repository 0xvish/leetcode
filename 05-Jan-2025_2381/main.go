package main

import "fmt"

func main() {
	s := "abc"
	shifts := [][]int{{0, 1, 1}, {1, 2, 0}}
	fmt.Println(shiftingLetters(s, shifts))
}

// ######## BRUTE FORCE SOLUTION ########
// func shiftingLetters(s string, shifts [][]int) string {
//     chars := []rune(s)
//     for _, shift := range shifts {
//         for j := shift[0] ; j <= shift[1] ; j ++ {
//             if shift[2] == 1 {
//                 chars[j] = 'a' + (chars[j]-'a'+1)%26
//             }

//             if shift[2] == 0 {
//                 chars[j] = 'a' + (chars[j]-'a'+25)%26
//             }
//         } 
//     }

//     return string(chars)
// }

func shiftingLetters(s string, shifts [][]int) string {
    n := len(s)
    shiftEffect := make([]int, n+1) // One extra element to handle range ending

    // Apply shift ranges
    for _, shift := range shifts {
        start, end, direction := shift[0], shift[1], shift[2]
        if direction == 1 {
            shiftEffect[start] += 1
            shiftEffect[end+1] -= 1
        } else {
            shiftEffect[start] -= 1
            shiftEffect[end+1] += 1
        }
    }

    // Compute prefix sum to determine the net shift at each index
    netShift := 0
    chars := []rune(s)
    for i := 0; i < n; i++ {
        netShift += shiftEffect[i]
        // Calculate final character with wrap-around using modular arithmetic
        chars[i] = 'a' + (chars[i]-'a'+rune((netShift%26+26)%26))%26
    }

    return string(chars)
}
