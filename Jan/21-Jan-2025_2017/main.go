package main

import (
	"fmt"
	"math"
)

func main() {
    grid := [][]int{{2,5,4},{1,5,1}}
    fmt.Println(gridGame(grid))
}

func gridGame(grid [][]int) int64 {
    N := len(grid[0])
    preSum1, preSum2 := make([]int64, N), make([]int64, N)

    preSum1[0], preSum2[0] = int64(grid[0][0]), int64(grid[1][0])
    for i := 1; i < N; i++ {
        preSum1[i] = preSum1[i-1] + int64(grid[0][i])
        preSum2[i] = preSum2[i-1] + int64(grid[1][i])
    }

    res := int64(math.MaxInt64)
    for i := 0; i < N; i++ {
        top := preSum1[N-1] - preSum1[i]
        bottom := int64(0)
        if i > 0 {
            bottom = preSum2[i-1]
        }

        secondBot := max(top, bottom)
        res = min(res, secondBot)
    }

    return res

}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}
