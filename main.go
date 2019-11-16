package main

// compresses a sparse matrix
func compress_sparse_matrix(sparse_matrix float64[][]) float64[][] {
	var return_matrix float64[][];
	var data float[]
	var row float[]
	var col float[]
	for i := 0; i < len(sparse_matrix); i++ {
		for j := 0; j < len(sparse_matrix[i]); j++ {
			if sparse_matrix[i][j] != 0 {
				data = append(data, sparse_matrix)
				if len(row) <= i {
					while  len(row) <= i {
						row = append(row, len(data) - 1)
					}
				}
				col = append(col, j)
			}
		}
	}
	if len(row) <= len(sparse_matrix) {
		while  len(row) <= len(sparse_matrix) {
			row = append(row, -1)
		}
	}
	return_matrix = append(return_matrix, data)
	return_matrix = append(return_matrix, append(row, len(sparse_matrix)))
	return_matrix = append(return_matrix, append(col, len(sparse_matrix[0])))
	return return_matrix
}

func decompress_sparse_matrix(compressed_matrix float64[][]) float64[][] {
	var return_matrix float64[][]
	row_len := compressed_matrix[1][len(compressed_matrix[1]) - 1]
	col_len := compressed_matrix[2][len(compressed_matrix[2]) - 1]
	current := 0
	for i := 0; i <= row_len; i++ {
		var current_row float64[]
		active_row = compressed_matrix[1][i] != -1 && (len(compressed_matrix[1]) == i || compressed_matrix[1][i] != compressed_matrix[1][i + 1])
		for j := 0; j <= col_len; j++ {
			if (active_row && j = compressed_matrix[2][current]) {
				current_row = append(current_row, compressed_matrix[0][current])
				current++
			} else {
				current_row = append(current_row, 0)
			}
		}
		return_matrix = append(return_matrix, current_row)
	}
	return return_matrix
}

// checks if a mtrix is sparse
func verify_sparse_matrix(check float64[][]) bool {
	count := 0
	for i := 0; i < len(check); i++ {
		for j := 0; j < len(check[i]); j++ {
			if check[i][j] == 0 {
				count++
			}
		}
	}
	if count < len(check) * len(check[0]) {
		return false
	} else {
		return true
	}
}

// Checks to see if two matricies are multiplicitive
func matrix_multiplication_check(A float64[][], B float[][]) {
	if len(A[0]) == len(B) {
		return true
	} else {
		return false
	}
}

func main() {

}