package main

type Config struct {
	Matrix_A_Loc            string        `json:"matrix_a"`
	Matrix_B_Loc            string        `json:"matrix_b"`
	MTX_File                bool          `json:"mtx_file"`
	Output_File             string        `json:"output_file_location"`
	Log_File                string        `json:"log_file_location"`
	Tolerance               float64       `json:"tolerance"`
	Max_Workers             int           `json:"max_workers"`
	Max_Iterations          int           `json:"max_iterations"`
	Max_Search_Directions   int           `json:"max_search_directions"`
}

func new_config() *Config {
	return &Config{
		Log_File:"./Sparse_Matrix_Operations.log",
		Output_File:"./Sparse_Matrix_Operations.txt",
	}
}
