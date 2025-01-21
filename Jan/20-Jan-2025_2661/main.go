package main

import "fmt"

func main() {
	arr := []int{12, 34, 67, 90, 45}
	mat := [][]int{
		{12, 34, 45, 67, 90},
		{45, 67, 90, 12, 34},
		{67, 90, 34, 12, 45},
		{90, 12, 67, 34, 45},
	}

	fmt.Println(firstCompleteIndex(arr, mat))
}

func firstCompleteIndex(arr []int, mat [][]int) int {
    if len(mat) == 0 || len(mat[0]) == 0 {
        return -1
    }

    getPosition := make(map[int][2]int)
    ROWS := len(mat)
    COLS := len(mat[0])
    for rowIndex, row := range mat {
        for colIndex, value := range row {
            getPosition[value] = [2]int{rowIndex, colIndex}
        }
    }

    rowCount := make([]int, ROWS)
    colCount := make([]int, COLS)

    for i, val := range arr {
        pos := getPosition[val]
        rowCount[pos[0]]++
        colCount[pos[1]]++

        if (colCount[pos[1]] == ROWS) || (rowCount[pos[0]] == COLS) {
            return i
        }
    }

    return -1

}