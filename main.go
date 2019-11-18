package main
//********************************************************************
//Created by: 	Joseph Jindrich
//Last update:	11/16/19
//********************************************************************

import(
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"io"
//	"io/ioutil"
	"encoding/json"
	"os"
	"strconv"
)

var config *Config = new_config()

// compresses a sparse matrix
func compress_sparse_matrix(sparse_matrix [][]float64) [][]float64 {
	var return_matrix [][]float64;
	var data []float64
	var row []float64
	var col []float64
	for i := 0; i < len(sparse_matrix); i++ {
		for j := 0; j < len(sparse_matrix[i]); j++ {
			if sparse_matrix[i][j] != 0 {
				data = append(data, sparse_matrix[i][j])
				if len(row) <= i {
					for len(row) <= i {
						row = append(row, float64(len(data) - 1))
					}
				}
				col = append(col, float64(j))
			}
		}
	}
	if len(row) < len(sparse_matrix) {
		for len(row) <= len(sparse_matrix) {
			row = append(row, -1)
		}
	}
	return_matrix = append(return_matrix, data)
	return_matrix = append(return_matrix, append(row, float64(len(sparse_matrix))))
	return_matrix = append(return_matrix, append(col, float64(len(sparse_matrix[0]))))
	return return_matrix
}

func decompress_sparse_matrix(compressed_matrix [][]float64) [][]float64 {
	var return_matrix [][]float64
	row_len := int(compressed_matrix[1][len(compressed_matrix[1]) - 1])
	col_len := int(compressed_matrix[2][len(compressed_matrix[2]) - 1])
	current := 0
	for i := 0; i < row_len; i++ {
		var current_row []float64
		active_row := compressed_matrix[1][i] != -1 && (len(compressed_matrix[1]) - 1 == i || compressed_matrix[1][i] != compressed_matrix[1][i + 1])
		for j := 0; j < col_len; j++ {
			if active_row && (i == row_len - 1 || compressed_matrix[1][i + 1] > float64(current)) && j == int(compressed_matrix[2][current]) {
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
func is_sparse_matrix(check [][]float64) bool {
	count := 0
	for i := 0; i < len(check); i++ {
		for j := 0; j < len(check[i]); j++ {
			if check[i][j] == 0 {
				count++
			}
		}
	}
	if float64(count) < float64((len(check) * len(check[0])) / 2) {
		return false
	} else {
		return true
	}
}

// Checks to see if two matricies are multiplicitive
func matrix_multiplication_check(A [][]float64, B [][]float64) bool {
	if len(A[0]) == len(B) {
		return true
	} else {
		return false
	}
}

func are_matricies_equal(A [][]float64, B [][]float64) bool {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		return false
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] != B[i][j] {
				return false
			}
		}
	}
	return true
}

//********************************************************************
// Name:	read_csv
// Description: This function reads any csv file passed in, and puts
//		it's data into an array of inputs. It also takes an 
//		int that it uses to only pull a input one over that
//		int times.
// Return:	returns an array of the type input.
//********************************************************************

func load_matrix_from_csv(file_location string) [][]float64 {
	var return_matrix [][]float64

	log.Print("Reading data file ", file_location)
	file, err := os.Open(file_location)
	if err != nil {
		log.Print("Error occured when opening ",
			file_location, "\n", err)
		os.Exit(-1)
	}
	reader := csv.NewReader(bufio.NewReader(file))
	//a for loop that continues until it reaches the end of the file.
	for {
		line, err := reader.Read()
		//error check for the end of a file.
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error occured while reading through ",
				    file_location + "\n\t\t", err)
			os.Exit(-1)
		}

		//parse through each data_entry and adds it to the data point.
		var new_row []float64
		for i := 0; i < len(line); i++ {
			data_entry, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				log.Print("Error occured while converting input on row ", i + 1 , " on line ",
					(len(return_matrix) + 1), " of the csv input file.\n\t\t", err)
				os.Exit(-1)
			}
			new_row = append(new_row,  data_entry)
		}
		return_matrix = append(return_matrix, new_row)
	}
	log.Print("Finished loading all training data from memory.")
	return return_matrix
}

func multiply_compressed_matricies(A [][]float64, B [][]float64) [][]float64 {
	A_col_len := int(A[2][len(A[2]) - 1])
	A_row_len := int(A[1][len(A[1]) - 1])
	B_row_len := int(B[1][len(B[1]) - 1])
	var product [][]float64
	for i := 0; i < A_col_len; i++ {
		var new_row []float64
		for j := 0; j < B_row_len; j++ {
			new_row = append(new_row, 0)
		}
		product = append(product, new_row)
	}
	a_row := 0
	for i := 0; i < len(A[0]); i++ {
		for (A[1][a_row] == -1 || i == int(A[1][a_row + 1])) && A_row_len != a_row {
			a_row++
		}
		for b_row := 0; b_row < B_row_len; b_row++ {
			for j := int(B[1][b_row]); j < len(B[0]) && j != -1 && (len(B[1]) - 1 == b_row || j < int(B[1][b_row + 1])); j++ {
				if a_row == int(B[2][j]) {
					product[b_row][int(A[2][i])] += A[0][i] * B[0][j]
				}
			}
		}
	}
	return product
}

//********************************************************************
// Name:	setup_log
// Description: This function sets up the log.
//********************************************************************

func setup_log () {
	if(config.Log_File != "") {
		log_file, err := os.OpenFile(config.Log_File, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("Error, Can not open ", config.Log_File , ": ", err)
		}
		log.SetOutput(log_file)
	}
}

func main() {
	var configPathFlag = flag.String("config", "./config.json", "path to configuration file")
	flag.Parse()
	if len(*configPathFlag) > 0 {
		file, err := os.Open(*configPathFlag)
		if err != nil {
			log.Fatal("Error, Can not access config: ", err)
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			log.Fatal("Error, Invalid config json: ", err)
		}
	}
	setup_log()
	log.Print("Starting Up")
	log.Print("Using config file ", *configPathFlag)

	A := load_matrix_from_csv("./Matrix a.csv")
	I := load_matrix_from_csv("./Matrix b.csv")
	cA := compress_sparse_matrix(A)
	cI := compress_sparse_matrix(I)
	fmt.Println(cA)
	fmt.Println(cI)
	C := multiply_compressed_matricies(cA, cI)
	fmt.Println(are_matricies_equal(A, C))
	for i := 0; i < len(A); i++ {
		fmt.Println(A[i])
		fmt.Println(C[i])
		fmt.Println()
	}
}
