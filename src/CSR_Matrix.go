package main

import(
	"fmt"
	"strconv"
)

type CSR_Matrix struct {
	row_len                int
	col_len                int
	data                   []float64
	row                    []int
	col                    []int
}

// compresses a sparse matrix
func compress_sparse_matrix(sparse_matrix [][]float64) CSR_Matrix {
	var csr CSR_Matrix
	csr.row_len = len(sparse_matrix)
	csr.col_len = len(sparse_matrix[0])
	for i := 0; i < csr.row_len; i++ {
		for j := 0; j < csr.col_len; j++ {
			if sparse_matrix[i][j] != 0 {
				csr.data = append(csr.data, sparse_matrix[i][j])
				for len(csr.row) <= i {
					csr.row = append(csr.row, len(csr.data) - 1)
				}
				csr.col = append(csr.col, j)
			}
		}
	}
	if len(csr.row) < csr.row_len {
		for len(csr.row) <= len(sparse_matrix) {
			csr.row = append(csr.row, -1)
		}
	}
	return csr
}

func (csr CSR_Matrix) decompress_sparse_matrix() [][]float64 {
	var return_matrix [][]float64
	current := 0
	for i := 0; i < csr.row_len; i++ {
		var current_row []float64
		active_row := csr.row[i] != -1 && (csr.row_len == i || csr.row[i] != csr.row[i + 1])
		for j := 0; j < csr.col_len; j++ {
			if active_row && (i == csr.row_len - 1 || csr.row[i + 1] > current) && j == csr.col[current] {
				current_row = append(current_row, csr.data[current])
				current++
			} else {
				current_row = append(current_row, 0)
			}
		}
		return_matrix = append(return_matrix, current_row)
	}
	return return_matrix
}

func (A CSR_Matrix) Times_CSR_matrix(B CSR_Matrix) [][]float64 {
	var product [][]float64
	for i := 0; i < A.col_len; i++ {
		var new_row []float64
		for j := 0; j < B.row_len; j++ {
			new_row = append(new_row, 0)
		}
		product = append(product, new_row)
	}
	a_row := 0
	for i := 0; i < len(A.data); i++ {
		for A.row_len != a_row + 1 && (A.row[a_row] == -1 || i == A.row[a_row + 1]) {
			a_row++
		}
		for b_row := 0; b_row < B.row_len; b_row++ {
			for j := B.row[b_row]; j < len(B.data) && j != -1 && (B.row_len == b_row || j < B.row[b_row + 1]); j++ {
				if a_row == B.col[j] {
					product[b_row][A.col[i]] += A.data[i] * B.data[j]
				}
			}
		}
	}
	return product
}

func (A CSR_Matrix) Times_vector(B []float64) []float64 {
	var product []float64
	for i := 0; i < A.col_len; i++ {
		product = append(product, 0)
	}
	a_row := 0
	for i := 0; i < len(A.data); i++ {
		for A.row_len != a_row + 1 && (A.row[a_row] == -1 || i == A.row[a_row + 1]) {
			a_row++
		}
		product[a_row] += A.data[i] * B[A.col[i]]
	}
	return product
}

func (A CSR_Matrix) Print() {
	fmt.Println("There are " + strconv.Itoa(A.row_len) + " rows and " + strconv.Itoa(A.col_len) + " columens in this CSR matrix.")
	fmt.Println("  data: ", A.data)
	fmt.Println("  row:  ", A.row)
	fmt.Println("  col:  ", A.col)
}
