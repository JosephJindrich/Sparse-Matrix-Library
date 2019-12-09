package main

import(
	"fmt"
	"strconv"
)

type CSR_Matrix struct {
	row_count                int
	col_count                int
	data                   []float64
	row                    []int
	col                    []int
}

// compresses a sparse matrix
func compress_sparse_matrix(sparse_matrix [][]float64) CSR_Matrix {
	var csr CSR_Matrix
	csr.row_count = len(sparse_matrix)
	csr.col_count = len(sparse_matrix[0])
	for i := 0; i < csr.row_count; i++ {
		for j := 0; j < csr.col_count; j++ {
			if sparse_matrix[i][j] != 0 {
				csr.data = append(csr.data, sparse_matrix[i][j])
				for len(csr.row) <= i {
					csr.row = append(csr.row, len(csr.data) - 1)
				}
				csr.col = append(csr.col, j)
			}
		}
	}
	if len(csr.row) < csr.row_count {
		for len(csr.row) <= len(sparse_matrix) {
			csr.row = append(csr.row, -1)
		}
	}
	return csr
}

func (csr CSR_Matrix) decompress() [][]float64 {
	var return_matrix [][]float64
	current := 0
	for i := 0; i < csr.row_count; i++ {
		var current_row []float64
		active_row := csr.row[i] != -1 && (csr.col_count == i + 1 || (i + 1 != csr.row_count && csr.row[i] != csr.row[i + 1]))
		for j := 0; j < csr.col_count; j++ {
			if active_row && (i + 1 == csr.col_count || csr.row[i + 1] > current) && j == csr.col[current] {
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

func (csr CSR_Matrix) Transpose() CSR_Matrix {
	var return_matrix CSR_Matrix
	var data_slices [][]float64
	var col_slices [][]int
	for i := 0; i < csr.col_count; i++ {
		var new_data_slice []float64
		var new_col_slice []int
		data_slices = append(data_slices, new_data_slice)
		col_slices = append(col_slices, new_col_slice)
	}
	current_row := 0
	for i := 0; i < len(csr.data); i++ {
		for(current_row + 1 != csr.row_count && csr.row[current_row + 1] <= i) {
			current_row = current_row + 1
		}
		data_slices[csr.col[i]] = append(data_slices[csr.col[i]], csr.data[i])
		col_slices[csr.col[i]] = append(col_slices[csr.col[i]], current_row)
	}
	for i := 0; i < csr.col_count; i++ {
		if len(data_slices[i]) > 0 {
			for len(return_matrix.row) <= i {
				return_matrix.row = append(return_matrix.row, len(return_matrix.data))
			}
			return_matrix.data = append(return_matrix.data, data_slices[i]...)
			return_matrix.col = append(return_matrix.col, col_slices[i]...)
		}
	}
	for len(return_matrix.row) < csr.col_count {
		return_matrix.row = append(return_matrix.row, -1)
	}
	return_matrix.row_count = csr.col_count
	return_matrix.col_count = csr.row_count
	return return_matrix
}

func (A CSR_Matrix) Times_CSR_matrix(B CSR_Matrix) [][]float64 {
	var product [][]float64
	for i := 0; i < A.col_count; i++ {
		var new_row []float64
		for j := 0; j < B.row_count; j++ {
			new_row = append(new_row, 0)
		}
		product = append(product, new_row)
	}
	a_row := 0
	for i := 0; i < len(A.data); i++ {
		for A.row_count != a_row + 1 && (A.row[a_row] == -1 || i == A.row[a_row + 1]) {
			a_row++
		}
		for b_row := 0; b_row < B.row_count; b_row++ {
			for j := B.row[b_row]; j < len(B.data) && j != -1 && (B.row_count == b_row + 1 || j < B.row[b_row + 1]); j++ {
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
	for i := 0; i < A.col_count; i++ {
		product = append(product, 0)
	}
	a_row := 0
	for i := 0; i < len(A.data); i++ {
		for A.row_count != a_row + 1 && (A.row[a_row] == -1 || i == A.row[a_row + 1]) {
			a_row++
		}
		product[a_row] += A.data[i] * B[A.col[i]]
	}
	return product
}

func (A CSR_Matrix) Print() {
	fmt.Println("There are " + strconv.Itoa(A.row_count) + " rows and " + strconv.Itoa(A.col_count) + " columens in this CSR matrix.")
	fmt.Println("  data:")
	for i := 1; i <= len(A.data); i++ {
		if i % 10 == 0 {
			fmt.Print(fmt.Sprintf("%.2f\n", A.data[i - 1]))
		} else {
			fmt.Print(fmt.Sprintf("%.2f, ", A.data[i - 1]))
		}
	}
	fmt.Println()
	fmt.Println("  row:")
	for i := 1; i <= len(A.row); i++ {
		if i % 10 == 0 {
			fmt.Print(fmt.Sprintf("%d\n", A.row[i - 1]))
		} else {
			fmt.Print(fmt.Sprintf("%d, ", A.row[i - 1]))
		}
	}
	fmt.Println()
	fmt.Println("  col:")
	for i := 1; i <= len(A.col); i++ {
		if i % 10 == 0 {
			fmt.Print(fmt.Sprintf("%d\n", A.col[i - 1]))
		} else {
			fmt.Print(fmt.Sprintf("%d, ", A.col[i - 1]))
		}
	}
	fmt.Println()
}
