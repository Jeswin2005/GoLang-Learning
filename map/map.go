package main

import "fmt"

func print(myMap map[int]string) {
	for key, value := range myMap {
		fmt.Printf("Key %d: %s\n", key, value)
	}
}

func demo(map1 map[int]string){
	map1[1] = "mango"
}

func main() {
	// Map defining

	myMap := make(map[int]string)
	// or var myMap[int][string]

	fmt.Println("Original map")
	myMap[1] = "Jason"
	myMap[2] = "kishore"
	myMap[3] = "jeswin"
	print(myMap)

	// check whether a key present or not
	name, ok := myMap[1]
	if ok {
		fmt.Println("id 1 : ", name)
	} else {
		fmt.Println("no id 1 present")
	}

	// delete an entry
	delete(myMap, 1)
	fmt.Println("Entry 1 deleted")
	fmt.Println("Updated map")
	print(myMap)

	// check whether a key present or not
	n, o := myMap[1]
	if o {
		fmt.Println("id 1 : ", n)
	} else {
		fmt.Println("no id 1 present")
	}


	map1 := make(map[int]string)
	map1[1] = "apple"
	map1[2] = "banana"

	demo(map1)

	fmt.Println(map1)
}
