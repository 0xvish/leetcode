package main

import (
	"fmt"
	"math/bits"
)

// minimizeXor calculates the minimum possible XOR value for num1 such that
// the number of set bits (1s) in the result matches the number of set bits in num2.
func minimizeXor(num1 int, num2 int) int {

    // Step 1: Count the number of set bits in num2
    num2_bits := bits.OnesCount(uint(num2))

    // Result variable to construct the desired number
    res := 0

    // Step 2: Use set bits of num1 from the highest position down to the lowest
    // Add these bits to `res` until `num2_bits` becomes zero
    for i := 31; i >= 0 && num2_bits > 0; i-- {
        // Check if the i-th bit in num1 is set
        if (num1 & (1 << i)) != 0 {
            // Set the i-th bit in `res`
            res |= (1 << i)
            // Decrement the count of remaining bits to set
            num2_bits--
        }
    }

    // Step 3: Add remaining set bits to `res` from the lowest position upward
    // This ensures the result has exactly the same number of set bits as `num2`
    for i := 0; i <= 31 && num2_bits > 0; i++ {
        // Check if the i-th bit in `res` is not already set
        if (res & (1 << i)) == 0 {
            // Set the i-th bit in `res`
            res |= (1 << i)
            // Decrement the count of remaining bits to set
            num2_bits--
        }
    }

    // Return the final result
    return res
}

func main() {
    // Test the minimizeXor function with num1 = 1 and num2 = 12
    // num1: 0001 (binary), num2: 1100 (binary), number of set bits in num2 = 2
    fmt.Println(minimizeXor(1, 12)) // Expected output: 3 (binary: 0011)
}

/*
Explanation of the logic:
1. **Objective**: Create a number `res` such that:
   - The number of set bits (1s) in `res` matches those in `num2`.
   - The XOR value between `num1` and `res` is minimized.

2. **Approach**:
   - Start with `num1` and try to retain its set bits in the highest positions
     (this minimizes the XOR value as higher bits contribute more to the value).
   - If more set bits are needed, fill them in the lowest unset positions of `res`.

3. **Why this works**:
   - Using higher bits of `num1` reduces the difference between `num1` and `res`.
   - Adding bits in the lowest positions minimizes the additional value introduced to `res`.
*/
