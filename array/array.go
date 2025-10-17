package main

import "fmt"

func main() {
	// 1. Array (Fixed Size)
	fruitsArray := [3]string{"apple", "banana", "orange"}

	fmt.Println("1. Array Elements (Fixed Size)")
	for i := 0; i < len(fruitsArray); i++ {
		fmt.Print(fruitsArray[i] + " ")
	}
	fmt.Println()

	// 2. Slice (Dynamic Size)
	vehiclesSlice := []string{"car", "bike", "auto"}

	fmt.Println("2. Slice Elements (Dynamic Size)")
	for _, vehicle := range vehiclesSlice {
		fmt.Print(vehicle + " ")
	}
	fmt.Println()

	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)

	// 3. 2D array
	fmt.Println("3. 2D Array")
	arr1 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr3 [3][3]int

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			arr3[i][j] = arr1[i][j] + arr2[i][j]
			fmt.Printf("%d", arr3[i][j])
		}
		fmt.Println()
	}

}
