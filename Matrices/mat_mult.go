package main

import(
	"fmt"
)




func main(){
	A:=fill_Mat(2,3,99)
	B:=fill_Mat(3,2,69)
	C:=Mat_Mult(A,B,len(A))
	fmt.Println("A",A)
	fmt.Println("B",B)
	fmt.Println("C",C)
}