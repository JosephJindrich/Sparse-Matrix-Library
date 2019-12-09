package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func  xxxxxxxx() {
	x := math.Sqrt(1)
	fmt.Println(x)
}

func new_vector(length int, random bool) []float64{
	var return_vector []float64
	for i := 0; i < length; i++ {
		if random {
			return_vector = append(return_vector, float64(rand.Int() % 10) - 5)
		} else {
			return_vector = append(return_vector, 0)
		}
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

func add_vectors(a []float64, b []float64) []float64{
	var return_vector []float64
	for i := 0; i < len(a); i++ {
		return_vector = append(return_vector, a[i] + b[i])
	}
	return return_vector
}

func multiply_vectors(a []float64, b []float64) float64{
	x := float64(0)
	for i := 0; i < len(a); i++ {
		x += a[i] * b[i]
	}
	return x
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

func calc_col_norm(A [][]float64, col int) float64 {
	norm := float64(0)
	for i:= 0; i < len(A); i++ {
		norm += A[i][col] * A[i][col]
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

func Matrix_copy(A [][]float64) [][]float64 {
	var B [][]float64
	for i := 0; i < len(A); i++ {
		var new_row []float64
		for j := 0; j < len(A[i]); j++{
			new_row = append(new_row, A[i][j])
		}
		B = append(B, new_row)
	}
	return B
}

func Vector_copy(a []float64) []float64 {
	var b []float64
	for i := 0; i < len(a); i++ {
		b = append(b, a[i])
	}
	return b
}

func calculate_QR(A [][]float64) ([][]float64, [][]float64) {
	Q := calculate_Q(A)
	R := multiply_matricies(Transpose(Q), A)
	for i := 1; i < len(R); i++ {
		for j := 0; j < i; j++ {
			R[i][j] = float64(0)
		}
	}
	return Q, R
}

func calculate_Q(A[][]float64) [][]float64 {
	V := Transpose(A)
	Q := copy_matrix(A)

	norm := calc_vector_norm(V[0])

	for i := 0; i < len(Q); i++ {
		Q[i][0] = Q[i][0] / norm
	}

	for i := 1; i < len(Q[0]); i++ {
		for j := 0; j < i; j++ {
			c := vector_times_column(V[i], Q, j)
			for k := 0; k < len(Q); k++ {
				Q[k][i] -= c * Q[k][j]
			}
		}
		norm = calc_col_norm(Q, i)
		for j := 0; j < len(Q); j++ {
			Q[j][i] = Q[j][i]/norm
		}
	}
	return Q
}

func copy_matrix(A [][]float64) [][]float64 {
	var B [][]float64
	for i := 0; i < len(A); i++ {
		var new_row []float64
		for j := 0; j < len(A[0]); j++ {
			new_row = append(new_row, A[i][j])
		}
		B = append(B, new_row)
	}
	return B
}

func Transpose(A [][]float64) [][]float64 {
	var A_transpose [][]float64
	for i := 0; i < len(A[0]); i++ {
		var new_row []float64
		A_transpose = append(A_transpose, append(new_row, A[0][i]))
	}
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			A_transpose[j] = append(A_transpose[j], A[i][j])
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

func matrix_times_vector(A [][]float64, b []float64) []float64 {
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

func find_alpha(R [][]float64, beta []float64) []float64 {
	var reverse_alpha []float64
	var alpha []float64
	for i := len(R) - 1; i >= 0; i-- {
		x := beta[i]
		for j := 0; j < len(reverse_alpha); j++ {
			x -= reverse_alpha[j] * R[i][(len(R) - 1) - j]
		}
		reverse_alpha = append(reverse_alpha, x / R[i][i])
	}
	for i := 0; i < len(reverse_alpha); i++ {
		alpha = append(alpha, reverse_alpha[len(reverse_alpha) - (i + 1)])
	}
	return alpha
}

func (A CSR_Matrix)GMRES(b []float64) []float64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	lowest_norm := float64(int(^uint(0) >> 1))
	var best_x []float64
	fail_count := 0
	Success := false
	iterations := 0
	worker_count := 0

	// Restart
	for iterations < config.Max_Iterations && !Success {
		wg.Add(1)
		go func() {
			fails := 0
			var x []float64
			defer wg.Done()
			var P [][]float64
			var B [][]float64
			x = new_vector(len(b), true)

			// r = b - Ax
			r := subtract_vectors(b, A.Times_vector(x))

			// initial residual = ||r||
			start_norm := calc_vector_norm(r)
			r_norm := start_norm
			prev := r_norm
			mux.Lock()
			if len(best_x) == 0 {
			lowest_norm = start_norm
			}
			mux.Unlock()

			// P = r / ||r||
			P = append(P, divide_vector(r, r_norm))

			// B = A * p
			B = append(B, A.Times_vector(P[0]))

			// loop
			for m := 1; m <= config.Max_Search_Directions; m++ {

				// B = QR
				Q, R := calculate_QR(Transpose(B))

				// Qt * r = beta
				beta := matrix_times_vector(Transpose(Q), r)

				// R * alpha = beta
				alpha := find_alpha(R, beta)

				// x = x + P * alpha
				x = add_vectors(x, matrix_times_vector(Transpose(P), alpha))

				// r = r - B * alpha
				r = subtract_vectors(r, matrix_times_vector(Transpose(B), alpha))
				//r = subtract_vectors(b, A.Times_vector(x))

				// residual norm = ||r||
				r_norm := calc_vector_norm(r)

				// check for convergence
				if r_norm <= config.Tolerance * start_norm {
					mux.Lock()
					Success = true
					fmt.Println(r_norm)
					best_x = x
					lowest_norm = r_norm
					iterations++
					mux.Unlock()
					break
				} else if Success {
					break
				}

				// p = r - c*p ...
				p := Vector_copy(r)
				for j := 0; j < m; j++ {
					// c = pt * r
					c := multiply_vectors(r, P[j])
					for k := 0; k < len(p); k++ {
						p[k] -= c * P[j][k]
					}
				}

				// p = p / ||p||
				p_norm := calc_vector_norm(p)
				if p_norm != 0 {
					p = divide_vector(p, p_norm)
				}

				// checking minimization
				if r_norm > prev {
					fails++
				}
				prev = r_norm

				//P = [P, p]
				P = append(P, p)

				// B = [B, A * p]
				B = append(B, A.Times_vector(p))
			}
			mux.Lock()
			fail_count += fails
			if (r_norm < lowest_norm) {
				best_x = x
				lowest_norm = r_norm
			}
			iterations++
			mux.Unlock()
		}()
		worker_count++
		if (worker_count >= config.Max_Workers && !Success) {
			fmt.Println("Finished setting up", config.Max_Workers, "and waiting for completion")
			wg.Wait()
			worker_count = 0;
			fmt.Println("All workers have finished,", iterations, "have been completed")
		}
	}
	wg.Wait()
	fmt.Println("Failed", fail_count, "times")
	if (!Success) {
		fmt.Println("Could not surpass Tolerence")
	}
	fmt.Println(lowest_norm)
	return best_x
}
