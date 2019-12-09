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
	"io/ioutil"
	"encoding/json"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var config *Config = new_config()

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
// Checks to see if two matricies are multiplicitive
func CSR_matrix_multiplication_check(A CSR_Matrix, B CSR_Matrix) bool {
	if A.col_count == B.row_count {
		return true
	} else {
		return false
	}
}

func are_matricies_equal(A [][]float64, B [][]float64) bool {
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A); i++ {
		if len(A[i]) != len(B[i]) {
			return false
		}
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] != B[i][j] {
				return false
			}
		}
	}
	return true
}

func append_vector(A [][]float64, b []float64) [][]float64{
	var return_vector [][]float64
	for i := 0; i < len(b); i++ {
		return_vector = append(return_vector, append(A[i], b[i]))
	}
	return return_vector
}

func transpose_vector(b []float64) [][]float64{
	var return_vector [][]float64
	for i := 0; i < len(b); i++ {
		var row []float64
		return_vector = append(return_vector, append(row, b[i]))
	}
	return return_vector
}


/*
func add_two_matricies(A [][]float64, B [][]float64) [][]float64 {
	row_count := int(A[1][len(A[1]) - 1])
	col_count := int(A[2][len(A[2]) - 1])
	var data []float64
	var row []float64
	var col []float64

	a_index := 0
	b_index := 0
	for len(A[0]) < a_index || len(B[0] < b_index {
		a_active_row := a[1][a_index] != -1 && (len(A[1]) - 1 == a_index || A[1][a_index] != A[1][a_index + 1])
		b_active_row := a[1][b_index] != -1 && (len(B[1]) - 1 == b_index || B[1][b_index] != B[1][b_index + 1])
		if a
	}
	row = append(row, row_count)
	col = append(col, col_count)
	return product
}
*/

//********************************************************************
// Name:	read_csv
// Description: This function reads any csv file passed in, and puts
//		it's data into an array of inputs. It also takes an 
//		int that it uses to only pull a input one over that
//		int times.
// Return:	returns an array of the type input.
//********************************************************************

func load_vector_from_csv(file_location string) []float64 {
	var return_vector []float64

	log.Print("Reading data file ", file_location)
	file, err := os.Open(file_location)
	if err != nil {
		log.Print("Error occured when opening ",
			file_location, "\n", err)
		os.Exit(-1)
	}
	reader := csv.NewReader(bufio.NewReader(file))
	//a for loop that continues until it reaches the end of the file.
	line, err := reader.Read()
	//error check for the end of a file.
	if err != nil {
		log.Println("Error occured while reading through ",
			    file_location + "\n\t\t", err)
		os.Exit(-1)
	}

	//parse through each data_entry and adds it to the data point.
	for i := 0; i < len(line); i++ {
		data_entry, err := strconv.ParseFloat(line[i], 64)
		if err != nil {
			log.Print("Error occured while converting input in the csv input file.\n\t\t", err)
			os.Exit(-1)
		}
		return_vector = append(return_vector,  data_entry)
	}
	log.Print("Finished loading all training data from memory.")
	return return_vector
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


func load_csr_from_mtx(file_location string) CSR_Matrix {
	var return_csr CSR_Matrix

	log.Print("Reading data file ", file_location)
	data, err := ioutil.ReadFile(file_location)
	if err != nil {
		log.Print("Error occured when opening ",
			file_location, "\n", err)
		os.Exit(-1)
	}
	lines := strings.Split(string(data), "\n")
	//a for loop that continues until it reaches the end of the file.
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 || string(lines[i][0]) == "%" {
			continue
		}
		line := strings.Split(lines[i], " ")
		col, _ := strconv.Atoi(line[0])
		row, _ := strconv.Atoi(line[1])
		value, _ := strconv.ParseFloat(line[2], 64)
		if (return_csr.row_count == 0) {
			return_csr.col_count = col
			return_csr.row_count = row
		} else {
			return_csr.data = append(return_csr.data, float64(int(value)))
			return_csr.col = append(return_csr.col, col - 1)
			for len(return_csr.row) < row {
				return_csr.row = append(return_csr.row, len(return_csr.data) - 1)
			}
		}
	}
	for len(return_csr.row) < return_csr.row_count {
		return_csr.row = append(return_csr.row, -1)
	}
	log.Print("Finished loading all training data from memory.")
	return return_csr
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
	log.Print("Starting Up")
}

func get_an_x_and_b(A CSR_Matrix) ([]float64, []float64) {
	var x []float64
	for i := 0; i < A.row_count; i++ {
		x = append(x, float64(rand.Int() % 10) - 5)
	}

	b := A.Times_vector(x)
	return x, b
}

func print_matrix(A [][]float64) {
	fmt.Println()
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			fmt.Print(A[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func calculate_error(experimental []float64, actual []float64) float64{
	e := math.Abs(calc_vector_norm(experimental))
	a := math.Abs(calc_vector_norm(actual))
	return (a - e) / (e + a)
}

func main() {
	// Read config from input
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
	log.Print("Using config file ", *configPathFlag)

	var cA CSR_Matrix
	if (config.MTX_File) {
		cA = load_csr_from_mtx(config.Matrix_A_Loc)
	} else {
		A := load_matrix_from_csv(config.Matrix_B_Loc)
		cA = compress_sparse_matrix(A)
	}
	actual_x, b := get_an_x_and_b(cA)
	fmt.Println(actual_x)
	x := cA.GMRES(b)
	fmt.Println(b)
	fmt.Println(cA.Times_vector(x))
	log.Print("Program has completed")
}
