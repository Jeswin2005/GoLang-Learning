package main

import "fmt"

func circumference [T int | float32](r T)T{
	return T(2*3.14*float32(r))
}

type Number interface{
	int | int8 | int32 | int64 | float32 | float64 | byte
}

func circumference_any [T Number](r T)T{
	return T(2*3.14*float64(r))
}

func print [T any](c T){
	fmt.Printf("The received value is of type %T and its value is %v",c,c)
}

func main(){
	var r1 int = 8
	var r2 float32 = 7.9

	fmt.Println("Radius(int) : ",circumference(r1))
	fmt.Println("Radius(float32) : ",circumference(r2))

	var r3 byte = 2
	fmt.Println("Radius(any type) : ",circumference_any(r3))

	print(4.5)
	fmt.Println()
	print("hi")

}