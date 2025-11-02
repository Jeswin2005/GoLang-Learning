package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var counter int = 0

func incremnet(wg *sync.WaitGroup){
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func main(){
	var wg sync.WaitGroup

	for i:=0 ; i<10 ;i++ {
		wg.Add(1)
		go incremnet(&wg)
	}

	wg.Wait()
	fmt.Println("Counter: ",counter)
}



