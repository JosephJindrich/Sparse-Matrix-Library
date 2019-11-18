package main

type input struct{
	values []float64
	target []float64
	position int
}

type Config struct {
	Matrix_A_Loc            string        `json:"matrix_a"`
	Matrix_B_Loc            string        `json:"matrix_b"`
	Output_File             string        `json:"output_file_location"`
	Log_File                string        `json:"log_file_location"`
}

func new_config() *Config {
	return &Config{
		Log_File:"./Sparse_Matrix_Operations.log",
		Output_File:"./Sparse_Matrix_Operations.txt",
	}
}
