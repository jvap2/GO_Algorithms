package main

import(
	"fmt"
)




func main(){
	A:=fill_Mat(2,3,99)
	B:=fill_Mat(3,2,69)
	C:=Mat_Mult(A,B,len(A),len(A[0]),len(B),len(B[0]))
	fmt.Println("A",A)
	fmt.Println("B",B)
	fmt.Println("C",C)


	D:=fill_Mat(4,4,420)
	fmt.Println("D", D)
	D=Inverse(D)
	fmt.Println("D^-1",D)
}