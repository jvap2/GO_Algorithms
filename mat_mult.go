package main
// import "fmt"
import(
	"math/rand"
	"fmt"
)



func fill(m int, n int, seed int64) [][]float64{
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

func main(){
	A:=fill(2,3,99)
	B:=fill(3,2,69)
	C:=Mat_Mult(A,B,len(A),len(A[0]),len(B),len(B[0]))
	fmt.Println("A",A)
	fmt.Println("B",B)
	fmt.Println("C",C)

}