package main

import(
	"math/rand"
	"fmt"
)


func fill_Mat(m int, n int, seed int64) [][]float64{
	r := rand.New(rand.NewSource(seed))
	mat:=make([][]float64,m)
	for i := range mat{
		mat[i]=make([]float64,n)
	}
	for i:=0; i<m;i++{
		for j:=0; j<n;j++{
			mat[i][j]=r.NormFloat64()
		}
	}
	return mat
}

func fill_Vect(m int, seed int64)[]float64{
	r := rand.New(rand.NewSource(seed))
	vect := make([]float64, m)
	for i:= range vect{
		vect[i]=r.NormFloat64()
	}
	return vect
}

func Mat_Mult(A [][]float64, B [][]float64, row_a int, col_a int, row_b int, col_b int) [][]float64{
	if col_a != row_b{
		fmt.Print("Dimensions are incompatible for matrix multiplication")
	}
	fSum:=0.0
	C:=make([][]float64, row_a)
	for i := range C{
		C[i]=make([] float64, col_b)
	}
	for i:=0; i<row_a;i++{
		for j:=0; j<col_b; j++{
			fSum=0
			for k:=0;k<col_a;k++{
				fSum+=A[i][k]*B[k][j]
			}
			C[i][j]=fSum
		}
	}
	return C
}

func Mat_Vect_Mult(A [][]float64, b []float64, row_a int, col_a int, v_len int) []float64{
	if col_a!=v_len{
		fmt.Print("Dimensions are incompatible")
	}
	out := make([]float64, row_a)
	var fSum float64
	for i:=0; i<row_a; i++{
		fSum=0.0
		for j:=0; j<col_a; j++{
			fSum+=A[i][j]*b[j]
		}
		out[i]=fSum
	}
	return out
}

func Dot_Product(A []float64, B []float64) float64 {
	if len(A)!=len(B){
		fmt.Print("Incompatible size vectors")
	}
	var out float64
	out=0.0
	for i:=0; i<len(A); i++{
		out+=A[i]*B[i]
	}
	return out
}
