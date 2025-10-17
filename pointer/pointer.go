package main

import "fmt"

func main() {

	// 1. pointer demo
	var a int
	var ptr1 *int
	var ptr2 **int

	a = 5
	ptr1 = &a
	ptr2 = &ptr1

	fmt.Println("value of a ", a)
	fmt.Println("value of ptr1 ", ptr1)
	fmt.Println("value of *ptr1 ", *ptr1)
	fmt.Println("value of ptr2 ", ptr2)
	fmt.Println("value of *ptr2 ", *ptr2)
	fmt.Println("value of **ptr2 ", **ptr2)

	// 2. Array of poiter
	var arr [3]*int

	x := 10
	arr[0] = &x

	y := 10
	arr[1] = &y

	z := 10
	arr[2] = &z
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
		fmt.Println(*arr[i])
		fmt.Println()
	}

	// 3. Pointer to a structure

	type Student struct {
		name string
		age  int
	}

	p := &Student{"Indu", 18}
	fmt.Println("Name=", p.name) // short hand dereferencing
	fmt.Println("Age=", (*p).age)
}
