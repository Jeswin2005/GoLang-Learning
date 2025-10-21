package main

import "fmt"

// Normal nested struct
type Address struct {
	City  string
	State string
}

// Outer struct with normal nested struct
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Outer struct with pointer to nested struct
type Employee struct {
	Name    string
	Age     int
	Address *Address
}

// Outer struct with embedded struct
type Student struct {
	Name string
	Age  int
	Address // embedded
}

func main() {
	// Normal nested struct
	p := Person{
		Name: "Jeswin",
		Age:  20,
		Address: Address{
			City:  "Chennai",
			State: "Tamil Nadu",
		},
	}
	fmt.Println("Normal Nested Struct:")
	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.Address.City)
	fmt.Println("State:", p.Address.State)
	fmt.Println()

	// Pointer to nested struct
	e := Employee{
		Name: "Jeswin",
		Age:  20,
		Address: &Address{
			City:  "Bangalore",
			State: "Karnataka",
		},
	}
	fmt.Println("Pointer to Nested Struct:")
	fmt.Println("Name:", e.Name)
	fmt.Println("City:", e.Address.City) // auto-dereference
	fmt.Println("State:", e.Address.State)
	fmt.Println()

	// Embedded struct
	s := Student{
		Name: "Jeswin",
		Age:  20,
		Address: Address{
			City:  "Mumbai",
			State: "Maharashtra",
		},
	}
	fmt.Println("Embedded Struct:")
	fmt.Println("Name:", s.Name)
	fmt.Println("City:", s.City)   // accessed directly
	fmt.Println("State:", s.State) // accessed directly
}
