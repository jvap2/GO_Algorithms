package main

type Toeplitz struct{
	r_size int
	c_size int
	row []float64
	col []float64
}

type Tridiagonal struct{
	size int
	l_diag []float64
	diag []float64
	u_diag []float64
}

type Vandermonde struct{
	size int
	col_1 []float64
}

type Identity struct{
	size int
}

type Diagonal struct{
	size int
	diag []float64
}

type Sparse_CSR struct{
	NNZ int
	V []float64
	COL_INDEX []int64
	ROW_INDEX []int64
}