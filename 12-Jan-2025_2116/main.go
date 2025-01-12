package main

import (
	"fmt"
)

func main() {
    fmt.Println(canBeValid("(()())", "000000"))
    fmt.Println(canBeValid("(()())", "001000"))
    fmt.Println(canBeValid("(()())", "000010"))
    fmt.Println(canBeValid("(()())", "000001"))
    fmt.Println(canBeValid("(()())", "000100"))
    fmt.Println(canBeValid("(()())", "000000"))
}

func canBeValid(s string, locked string) bool {

    locked_stack := []int{}
    unlocked_stack := []int{}

    if len(s)%2 != 0 {
        return false
    }

    for i:=0 ; i<len(s) ; i++ {
        if locked[i] == '0' {
            unlocked_stack = append(unlocked_stack, i)
        } else if s[i] == '(' {
            locked_stack = append(locked_stack, i)
        } else {
            if len(locked_stack) > 0 {
                locked_stack = locked_stack[:len(locked_stack)-1]
            } else if len(unlocked_stack) > 0 {
                unlocked_stack = unlocked_stack[:len(unlocked_stack)-1]
            } else {
                return false
            }
        }
    }

    for len(locked_stack)>0 && len(unlocked_stack)>0 && locked_stack[len(locked_stack)-1] < unlocked_stack[len(unlocked_stack)-1] {
        locked_stack = locked_stack[:len(locked_stack)-1]
        unlocked_stack = unlocked_stack[:len(unlocked_stack)-1]
    }

    if (len(locked_stack)>0) {
        return false
    }

    if len(unlocked_stack)%2 != 0 {
        return false
    }

    return true

}