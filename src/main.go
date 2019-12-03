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
	if A.col_len == B.row_len {
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
		return_vector = append( return_vector, append(A[i], b[i]))
	}
	return return_vector
}


/*
func add_two_matricies(A [][]float64, B [][]float64) [][]float64 {
	row_len := int(A[1][len(A[1]) - 1])
	col_len := int(A[2][len(A[2]) - 1])
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
	row = append(row, row_len)
	col = append(col, col_len)
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

	A := load_matrix_from_csv(config.Matrix_A_Loc)
	b := load_vector_from_csv(config.Matrix_B_Loc)
	cA := compress_sparse_matrix(A)

	cA.Print()
	fmt.Println(b)

	Z := cA.Times_vector(b)

	fmt.Println(Z)
}
