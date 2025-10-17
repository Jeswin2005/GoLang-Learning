package main

import "fmt"

type Engine struct {
    Power int
}

func (e Engine) Start() string {
    return fmt.Sprintf("Engine started with %d HP", e.Power)
}

// Car embeds Engine
type Car struct {
    Engine 
    Model  string
}

func main() {
    c := Car{
        Engine: Engine{Power: 300},
        Model:  "Mustang",
    }

    fmt.Println("Car Power:", c.Power) 
    fmt.Println("Car Model:", c.Model) 

    fmt.Println(c.Start()) 
}