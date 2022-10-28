package q10_valid_sudoku

import "fmt"

// validSudokuVersion2...
func validSudokuVersion2(sudoku [][]byte) bool {
	var numberRow = make(map[int]map[byte]int)
	var numberCol = make(map[int]map[byte]int)
	var numberCube = make(map[string]map[byte]int)
	for r, row := range sudoku {
		for c, node := range row {
			if node == byte('.') {
				continue
			} else {
				if numberRow[r] != nil && numberRow[r][node] != 0 {
					fmt.Println("row", r, c, node)
					return false
				}
				if numberRow[r] == nil {
					numberRow[r] = map[byte]int{node: 1}
				} else {
					numberRow[r][node] = 1
				}
				//numberRow[r][node] = 1
				if numberCol[c] != nil && numberCol[c][node] != 0 {
					fmt.Println("col", r, c, node)
					return false
				}
				if numberCol[c] == nil {
					numberCol[c] = map[byte]int{node: 1}
				} else {
					numberCol[c][node] = 1
				}
				rx := r / 3
				cx := c / 3
				key := fmt.Sprintf("%d-%d", rx, cx)
				if numberCube[key] != nil && numberCube[key][node] != 0 {
					fmt.Println("cube", r, c, node)
					return false
				}
				if numberCube[key] == nil {
					numberCube[key] = map[byte]int{node: 1}
				} else {
					numberCube[key][node] = 1
				}
			}
		}
	}
	return true
}

//func validSudoku(sudoku [][]byte) (valid bool) {
//	var numberRow []map[byte]int
//	var numberCol []map[byte]int
//	var numberCube map[string]map[byte]int
//numberMap := map[string]int{}
//numberRow := map[string]int{}
//blankByte := []byte(".")[0]
//blankByte := "."
//numberCube := map[string]map[string]int{}
// 时间换空间，
//a := 'a'
//fmt.Println(reflect.TypeOf()ypeof(a))
//r, c := 0, 0
//for r < 9 {
//	for c < 9 {
//		if _, ok := numberMap[sudoku[r][c]]; ok {
//			valid = false
//			return
//		}
//		if sudoku[r][c] != blankByte {
//			numberMap[sudoku[r][c]] = 1
//		}
//		c += 1
//	}
//	r += 1
//	numberMap = map[string]int{}
//}
//i, j := 0, 0
//numberMap = map[string]int{}
//for i < 9 {
//	for j < 9 {
//		if _, ok := numberMap[sudoku[j][i]]; ok {
//			valid = false
//			return
//		}
//		if sudoku[j][i] != "." {
//			numberMap[sudoku[j][i]] = 1
//		}
//		j += 1
//	}
//	i += 1
//	numberMap = map[string]int{}
//}
//small area 3 x 3
//
//valid = true
//return
//}
