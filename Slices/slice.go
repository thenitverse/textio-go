package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range rows {
		matrix[i] = make([]int, cols)
		for j := range cols {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func main() {
	fmt.Println(createMatrix(3, 3))
	fmt.Println(createMatrix(4, 4))
	fmt.Println(createMatrix(0, 0))
}

/*output:
[[0 0 0] [0 1 2] [0 2 4]]
[[0 0 0 0] [0 1 2 3] [0 2 4 6] [0 3 6 9]]
[]
*/
