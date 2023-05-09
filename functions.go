package main

import(
	"math/rand"
	"fmt"
)

func CGS(A [][]float64, b []float64, x []float64)[]float64{
	A_T:=Tranpose(A)
	r:=Mat_Vect_Mult(A_T,Vect_Add(b,Mat_Vect_Mult(A,x, len(A),len(A[0]), len(x)),false),len(A_T),len(A_T[0]),len(b))
	d:=r
	check:=true
	tol:=1e-9
	for check{
		q:=Mat_Vect_Mult(A_T,Mat_Vect_Mult(A,d,len(A),len(A[0]),len(d)),len(A_T),len(A_T[0]),len(A))
		Ad:=Mat_Vect_Mult(A,d,len(A),len(A[0]),len(d))
		lambda:=Dot_Product(r,r)/Dot_Product(Ad,Ad)
		if lambda<=tol{
			break
		}
		x=Vect_Add(x,Const_Vect_Mult(lambda,d),true)
		r_new:=Vect_Add(r,Const_Vect_Mult(lambda,q),false)
		beta:=Dot_Product(r_new,r_new)/Dot_Product(r,r)
		r=r_new
		d=Vect_Add(r,Const_Vect_Mult(beta,d),true)
	}
	return x
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

func Const_Vect_Mult(q float64, B []float64)[]float64{
	out:=make([]float64, len(B))
	for i:=0; i<len(B);i++{
		out[i]=q*B[i]
	}
	return out
} 


func Vect_Add(A []float64, B []float64, sign bool) []float64{
	if len(A)!=len(B){
		fmt.Print("Vector sizes are incompatible")
	}
	out:=make([]float64, len(A))
	if sign{
		for i:=0; i<len(A); i++{
			out[i]=A[i]+B[i]
		}
	}else{
		for i:=0; i<len(A); i++{
			out[i]=A[i]-B[i]
		}
	}
	return out
}

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

func Tranpose(A [][]float64) [][]float64{
	A_T:=make([][]float64, len(A[0]))
	for i:= range(A_T){
		A_T[i]=make([]float64, len(A))
	}
	for i:=0; i<len(A);i++{
		for j:=0; j<len(A[0]); j++{
			A_T[j][i]=A[i][j]
		} 
	}
	return A_T
}