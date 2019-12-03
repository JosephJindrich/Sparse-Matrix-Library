package main

import (
	"math"
)

func new_vector(length int, random bool) []float64{
	var return_vector []float64
	for i := 0; i < length; i++ {
		return_vector = append(return_vector, 0)
	}
	return return_vector
}

func subtract_vectors(a []float64, b []float64) []float64{
	var return_vector []float64
	for i := 0; i < len(a); i++ {
		return_vector = append(return_vector, a[i] - b[i])
	}
	return return_vector
}

func divide_vector(a []float64, x float64) []float64{
	var return_vector []float64
	for i := 0; i < len(a); i++ {
		return_vector = append(return_vector, a[i]/x)
	}
	return return_vector
}

func calc_vector_norm(a []float64) float64 {
	norm := float64(0)
	for i:= 0; i < len(a); i++ {
		norm += a[i] * a[i]
	}
	return math.Sqrt(norm)
}

func calc_col_norm(a [][]float64, col int) float64 {
	norm := float64(0)
	for i:= 0; i < len(a); i++ {
		norm += a[i][col] * a[i][col]
	}
	return math.Sqrt(norm)
}

func vector_times_column(a []float64, B [][]float64, col int) float64 {
	x := float64(0)
	for i := 0; i < len(a); i++ {
		x += a[i] * B[i][col]
	}
	return x
}

func calculate_Q(A[][]float64) [][]float64 {
	V := Transpose(A)
	Q := A

	norm := calc_vector_norm(V[0])
	for i := 0; i < len(Q); i++ {
		Q[i][0] = Q[i][0]/norm
	}

	for i := 1; i < len(V); i++ {
		for j := 0; j < i; j++ {
			c := vector_times_column(V[i], Q, j)
			for k := 0; k < len(Q); k++ {
				Q[i][k] -= c * Q[j][k]
			}
		}
		norm = calc_col_norm(Q, i)
		for j := 0; j < len(Q); j++ {
			Q[j][i] = Q[j][i]/norm
		}
	}
	return Q
}

func Transpose(A [][]float64) [][]float64 {
	var A_transpose [][]float64
	for i := 0; i < len(A); i++ {
		if i == 0 {
			var new_row []float64
			for j := 0; j < len(A[0]); j++ {
				A_transpose = append(A_transpose, append(new_row, A[i][j]))
			}
		} else {
			for j := 0; j < len(A[0]); j++ {
				A_transpose[i] = append(A_transpose[i], A[i][j])
			}
		}
	}
	return A_transpose
}

func multiply_matricies(A [][]float64, B [][]float64) [][]float64{
	var C [][]float64
	for i := 0; i < len(A); i++ {
		var new_row []float64
		for j := 0; j < len(B[0]); j++ {
			new_row = append(new_row, 0)
			for k := 0; k < len(B); k++ {
				new_row[j] += A[i][k] * B[k][j]
			}
		}
		C = append(C, new_row)
	}
	return C
}

func matrix_times_vector(A [][]float64, b []float64) []float64{
	var c []float64
	for i := 0; i < len(A); i++ {
		x := float64(0)
		for j := 0; j < len(b); j++ {
			x += A[i][j] * b[j]
		}
		c = append(c, x)
	}
	return c
}

func calculate_QR(A [][]float64) ([][]float64, [][]float64) {
	Q := calculate_Q(A)
	R := multiply_matricies(Transpose(Q), A)
	return Q, R
}

func GMRES(A CSR_Matrix, b []float64) {
	var P [][]float64
	var B [][]float64

	x := new_vector(len(b), false)
	r := subtract_vectors(b, A.Times_vector(x))
	start_norm := calc_vector_norm(r)
	r_norm := start_norm

	for m := 1; m < 10; m++ {
		P = append_vector(P, divide_vector(r, r_norm))
		B = append_vector(B, r)
		Q, R := calculate_QR(B)
		beta := matrix_times_vector(Transpose(Q), r)
		alpha := matrix_times_vector(Transpose(R), beta)
		x := add_vectors(x, matrix_times_vector(P, alpha)
		r := subtract_vectors(r, matrix_times_vector(B, alpha)
		r_norm := calc_vector_norm(r)
		if r_norm < config.Tolerence * start_norm {
			return x
		}
		p := r
		for j := 0; j < m; j++ {
			p -= 
		}
		p = divide_vector(p, valc_vector_norm(p))
		P = append(P, p)
		B = append(B, A.times_vector(p))
	}
}
